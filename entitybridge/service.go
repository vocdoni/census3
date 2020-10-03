package entitybridge

import (
	"context"

	"github.com/vocdoni/tokenstate"
)

type EntityBridgeService struct {
	Web3 *tokenstate.Web3
}

func NewEntityBridgeService() *EntityBridgeService {
	return &EntityBridgeService{
		Web3: new(tokenstate.Web3),
	}
}
func (bs *EntityBridgeService) Init(ctx context.Context, web3Endpoint, contractAddress string) error {
	// conect to eth and the contract
	if err := bs.Web3.Init(ctx, web3Endpoint, contractAddress); err != nil {
		return err
	}
	return nil
}
