package api

import (
	"context"
	"math/big"
	"time"

	queries "github.com/vocdoni/census3/db/sqlc"
)

// createDummyStrategy creates the default strategy for a given token. This
// basic strategy only includes the holders of the given token which have a
// balance positive balance (holder_balance > 0).
//
// TODO: Only for the MVP, remove it.
func (capi *census3API) createDummyStrategy(tokenID []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := capi.sqlc.CreateStategy(ctx, "test")
	if err != nil {
		return err
	}
	strategyID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	_, err = capi.sqlc.CreateStrategyToken(ctx, queries.CreateStrategyTokenParams{
		StrategyID: strategyID,
		TokenID:    tokenID,
		MinBalance: big.NewInt(0).Bytes(),
		MethodHash: []byte("test"),
	})
	return err
}
