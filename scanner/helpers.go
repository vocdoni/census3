package scanner

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/db"
	"github.com/vocdoni/census3/db/annotations"
	queries "github.com/vocdoni/census3/db/sqlc"
	"go.vocdoni.io/dvote/log"
)

// SaveHolders saves the given holders in the database. It updates the token
// synced status if it is different from the received one. Then, it creates,
// updates or deletes the token holders in the database depending on the
// calculated balance.
// WARNING: the following code could produce holders with negative balances
// in the database. This is because the scanner does not know if the token
// holder is a contract or not, so it does not know if the balance is
// correct or not. The scanner assumes that the balance is correct and
// updates it in the database:
//  1. To get the correct holders from the database you must filter the
//     holders with negative balances.
//  2. To get the correct balances you must use the contract methods to get
//     the balances of the holders.
func SaveHolders(db *db.DB, ctx context.Context, token ScannerToken,
	holders map[common.Address]*big.Int, newTransfers, lastBlock uint64,
	synced bool, totalSupply *big.Int,
) (int, int, error) {
	// create a tx to use it in the following queries
	tx, err := db.RW.BeginTx(ctx, nil)
	if err != nil {
		return 0, 0, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
			log.Errorf("error rolling back tx: %v, token=%s chainID=%d externalID=%s",
				err, token.Address.Hex(), token.ChainID, token.ExternalID)
		}
	}()
	qtx := db.QueriesRW.WithTx(tx)
	// create, update or delete token holders
	created, updated := 0, 0
	for addr, balance := range holders {
		// get the current token holder from the database
		currentTokenHolder, err := qtx.GetTokenHolderEvenZero(ctx, queries.GetTokenHolderEvenZeroParams{
			TokenID:    token.Address.Bytes(),
			ChainID:    token.ChainID,
			ExternalID: token.ExternalID,
			HolderID:   addr.Bytes(),
		})
		if err != nil {
			if !errors.Is(sql.ErrNoRows, err) {
				return created, updated, err
			}
			// if the token holder not exists, create it
			_, err = qtx.CreateTokenHolder(ctx, queries.CreateTokenHolderParams{
				TokenID:    token.Address.Bytes(),
				ChainID:    token.ChainID,
				ExternalID: token.ExternalID,
				HolderID:   addr.Bytes(),
				BlockID:    lastBlock,
				Balance:    balance.String(),
			})
			if err != nil {
				return created, updated, err
			}
			created++
			continue
		}
		// parse the current balance of the holder
		currentBalance, ok := new(big.Int).SetString(currentTokenHolder.Balance, 10)
		if !ok {
			return created, updated, fmt.Errorf("error parsing current token holder balance")
		}
		// if both balances are zero, continue with the next holder to prevent
		// UNIQUES constraint errors
		if balance.Cmp(big.NewInt(0)) == 0 && currentBalance.Cmp(big.NewInt(0)) == 0 {
			continue
		}
		// calculate the new balance of the holder by adding the current balance
		// and the new balance
		newBalance := new(big.Int).Add(currentBalance, balance)
		// update the token holder in the database with the new balance.
		// WANING: the balance could be negative so you must filter the holders
		// with negative balances to get the correct holders from the database.
		_, err = qtx.UpdateTokenHolderBalance(ctx, queries.UpdateTokenHolderBalanceParams{
			TokenID:    token.Address.Bytes(),
			ChainID:    token.ChainID,
			ExternalID: token.ExternalID,
			HolderID:   addr.Bytes(),
			BlockID:    currentTokenHolder.BlockID,
			NewBlockID: lastBlock,
			Balance:    newBalance.String(),
		})
		if err != nil {
			return created, updated, fmt.Errorf("error updating token holder: %w", err)
		}
		updated++
	}
	// get the token info from the database to update ir
	tokenInfo, err := qtx.GetToken(ctx,
		queries.GetTokenParams{
			ID:         token.Address.Bytes(),
			ChainID:    token.ChainID,
			ExternalID: token.ExternalID,
		})
	if err != nil {
		return created, updated, err
	}
	// update the synced status, last block, the number of analysed transfers
	// (for debug) and the total supply in the database
	_, err = qtx.UpdateTokenStatus(ctx, queries.UpdateTokenStatusParams{
		ID:                token.Address.Bytes(),
		ChainID:           token.ChainID,
		ExternalID:        token.ExternalID,
		Synced:            synced,
		LastBlock:         int64(lastBlock),
		AnalysedTransfers: tokenInfo.AnalysedTransfers + int64(newTransfers),
		TotalSupply:       annotations.BigInt(totalSupply.String()),
	})
	if err != nil {
		return created, updated, err
	}
	// close the database tx and commit it
	if err := tx.Commit(); err != nil {
		return created, updated, err
	}
	return created, updated, nil
}
