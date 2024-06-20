package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/scanner/providers/web3"
	"go.vocdoni.io/dvote/httprouter"
	api "go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

func (capi *census3API) initHoldersHandlers() error {
	if err := capi.endpoint.RegisterMethod("/holders/{tokenID}", "GET",
		api.MethodAccessTypeAdmin, capi.launchHoldersAtLastBlock); err != nil {
		return err
	}
	return capi.endpoint.RegisterMethod("/holders/queue/{queueID}", "GET",
		api.MethodAccessTypePublic, capi.enqueueHoldersAtLastBlock)
}

func (capi *census3API) launchHoldersAtLastBlock(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// get contract address from the tokenID query param and decode check if
	// it is provided, if not return an error
	strAddress := ctx.URLParam("tokenID")
	if strAddress == "" {
		return ErrMalformedToken.With("tokenID is required")
	}
	address := common.HexToAddress(strAddress)
	// get chainID from query params and decode it as integer, if it's not
	// provided or it's not a valid integer return an error
	strChainID := ctx.Request.URL.Query().Get("chainID")
	if strChainID == "" {
		return ErrMalformedChainID.With("chainID is required")
	}
	chainID, err := strconv.Atoi(strChainID)
	if err != nil {
		return ErrMalformedChainID.WithErr(err)
	} else if chainID < 0 {
		return ErrMalformedChainID.With("chainID must be a positive number")
	}
	// get externalID from query params and decode it as string, it is optional
	// so if it's not provided continue
	externalID := ctx.Request.URL.Query().Get("externalID")
	// list holders balances at last block in background
	queueID := capi.queue.Enqueue()
	go func() {
		balances, lastBlockNumber, err := capi.listHoldersAtLastBlock(address, uint64(chainID), externalID)
		if err != nil {
			if ok := capi.queue.Fail(queueID, err); !ok {
				log.Errorf("error updating list holders at block queue %s", queueID)
			}
			return
		}
		queueData := map[string]any{
			"size":    len(balances),
			"block":   lastBlockNumber,
			"holders": balances,
		}
		if ok := capi.queue.Done(queueID, queueData); !ok {
			log.Errorf("error updating list holders at block queue %s", queueID)
		}
	}()
	// encode and send the queueID
	res, err := json.Marshal(QueueResponse{QueueID: queueID})
	if err != nil {
		return ErrEncodeStrategy.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}

func (capi *census3API) listHoldersAtLastBlock(address common.Address,
	chainID uint64, externalID string,
) (map[string]string, uint64, error) {
	// get token information from the database
	internalCtx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()
	tokenData, err := capi.db.QueriesRO.GetToken(internalCtx,
		queries.GetTokenParams{
			ID:         address.Bytes(),
			ChainID:    chainID,
			ExternalID: externalID,
		})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, 0, ErrNotFoundToken.WithErr(err)
		}
		return nil, 0, ErrCantGetToken.WithErr(err)
	}
	// get token holders count
	holders, err := capi.db.QueriesRO.ListTokenHolders(internalCtx,
		queries.ListTokenHoldersParams{
			TokenID:    address.Bytes(),
			ChainID:    chainID,
			ExternalID: externalID,
		})
	if err != nil {
		return nil, 0, ErrCantGetTokenHolders.WithErr(err)
	}
	// if the token is external, return an error
	provider, err := capi.holderProviders.GetProvider(internalCtx, tokenData.TypeID)
	if err != nil {
		return nil, 0, ErrCantCreateCensus.WithErr(fmt.Errorf("token type not supported: %w", err))
	}
	if provider.IsExternal() {
		return nil, 0, ErrCantCreateCensus.With("not implemented for external providers")
	}
	if err := provider.SetRef(web3.Web3ProviderRef{
		HexAddress: common.Bytes2Hex(tokenData.ID),
		ChainID:    tokenData.ChainID,
	}); err != nil {
		return nil, 0, ErrInitializingWeb3.WithErr(err)
	}

	// get last block of the network
	lastBlockNumber, err := provider.LatestBlockNumber(internalCtx, nil)
	if err != nil {
		return nil, 0, ErrCantGetLastBlockNumber.WithErr(err)
	}
	// get holders balances at last block
	balances := make(map[string]string)
	for i, holder := range holders {
		log.Infow("getting balance",
			"holder", common.BytesToAddress(holder.HolderID).String(),
			"token", address.String(),
			"progress", fmt.Sprintf("%d/%d", i+1, len(holders)))
		holderAddress := common.BytesToAddress(holder.HolderID)
		balance, err := provider.BalanceAt(internalCtx, holderAddress, nil, lastBlockNumber)
		if err != nil {
			return nil, lastBlockNumber, ErrCantGetTokenHolders.WithErr(err)
		}
		if balance.Cmp(big.NewInt(0)) > 0 {
			balances[holderAddress.String()] = balance.String()
		}
	}
	return balances, lastBlockNumber, nil
}

func (capi *census3API) enqueueHoldersAtLastBlock(msg *api.APIdata, ctx *httprouter.HTTPContext) error {
	// parse queueID from url
	queueID := ctx.URLParam("queueID")
	if queueID == "" {
		return ErrMalformedStrategyQueueID
	}
	// try to get and check if the strategy is in the queue
	queueItem, exists := capi.queue.IsDone(queueID)
	if !exists {
		return ErrNotFoundStrategy.Withf("the ID %s does not exist in the queue", queueID)
	}
	// check if it is not finished or some error occurred
	if queueItem.Done && queueItem.Error == nil {
		// remove the item from the queue
		capi.queue.Dequeue(queueID)
	}
	// encode item response and send it
	res, err := json.Marshal(queueItem)
	if err != nil {
		return ErrEncodeQueueItem.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}
