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
	"github.com/vocdoni/census3/service/web3"
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
			if ok := capi.queue.Update(queueID, true, nil, err); !ok {
				log.Errorf("error updating list holders at block queue %s", queueID)
			}
			return
		}
		queueData := map[string]any{
			"size":    len(balances),
			"block":   lastBlockNumber,
			"holders": balances,
		}
		if ok := capi.queue.Update(queueID, true, queueData, nil); !ok {
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
	tokenData, err := capi.db.QueriesRO.TokenByIDAndChainIDAndExternalID(internalCtx,
		queries.TokenByIDAndChainIDAndExternalIDParams{
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
	holders, err := capi.db.QueriesRO.TokenHoldersByTokenIDAndChainIDAndExternalID(internalCtx,
		queries.TokenHoldersByTokenIDAndChainIDAndExternalIDParams{
			TokenID:    address.Bytes(),
			ChainID:    chainID,
			ExternalID: externalID,
		})
	if err != nil {
		return nil, 0, ErrCantGetTokenHolders.WithErr(err)
	}
	// if the token is external, return an error
	// TODO: implement external token holders
	if _, isExternal := capi.extProviders[tokenData.TypeID]; isExternal {
		return nil, 0, ErrCantCreateCensus.With("not implemented for external providers")
	}
	// get correct web3 uri provider
	w3URI, exists := capi.w3p.EndpointByChainID(tokenData.ChainID)
	if !exists {
		return nil, 0, ErrChainIDNotSupported.With("chain ID not supported")
	}
	w3, err := w3URI.GetClient(web3.DefaultMaxWeb3ClientRetries)
	if err != nil {
		return nil, 0, ErrInitializingWeb3.WithErr(err)
	}
	// get last block of the network
	lastBlockNumber, err := w3.BlockNumber(internalCtx)
	if err != nil {
		return nil, 0, ErrCantGetLastBlockNumber.WithErr(err)
	}
	bLastBlockNumber := new(big.Int).SetUint64(lastBlockNumber)
	// get holders balances at last block
	balances := make(map[string]string)
	for i, holder := range holders {
		log.Infow("getting balance",
			"holder", common.BytesToAddress(holder.ID).String(),
			"token", address.String(),
			"progress", fmt.Sprintf("%d/%d", i+1, len(holders)))
		holderAddress := common.BytesToAddress(holder.ID)
		balance, err := w3.BalanceAt(internalCtx, holderAddress, bLastBlockNumber)
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
	exists, done, data, err := capi.queue.Done(queueID)
	if !exists {
		return ErrNotFoundStrategy.Withf("the ID %s does not exist in the queue", queueID)
	}
	// init the queue response
	queueStrategy := GetHoldersAtLastBlockResponse{
		Done:  done,
		Error: err,
	}
	// check if it is not finished or some error occurred
	if done && err == nil {
		queueStrategy.HoldersAtBlock = &TokenHoldersAtBlock{
			Size:        data["size"].(int),
			BlockNumber: data["block"].(uint64),
			Holders:     data["holders"].(map[string]string),
		}
		// remove the item from the queue
		capi.queue.Dequeue(queueID)
	}
	// encode item response and send it
	res, err := json.Marshal(queueStrategy)
	if err != nil {
		return ErrEncodeQueueItem.WithErr(err)
	}
	return ctx.Send(res, api.HTTPstatusOK)
}
