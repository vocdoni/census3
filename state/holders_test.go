package state

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
)

func TestTokenHoldersInit(t *testing.T) {
	c := qt.New(t)
	th := new(TokenHolders).Init(MonkeysAddress, CONTRACT_TYPE_ERC20, MonkeysCreationBlock, 0, "")
	c.Assert(th.address.String(), qt.Equals, MonkeysAddress.String())
	c.Assert(th.ctype, qt.Equals, CONTRACT_TYPE_ERC20)
	c.Assert(th.lastBlock.Load(), qt.Equals, MonkeysCreationBlock)
	c.Assert(th.synced.Load(), qt.IsFalse)
}

func TestHolders(t *testing.T) {
	c := qt.New(t)
	th := new(TokenHolders).Init(MonkeysAddress, CONTRACT_TYPE_ERC20, MonkeysCreationBlock, 5, "")
	c.Assert(th.address.String(), qt.Equals, MonkeysAddress.String())
	c.Assert(th.ctype, qt.Equals, CONTRACT_TYPE_ERC20)
	c.Assert(th.lastBlock.Load(), qt.Equals, MonkeysCreationBlock)
	c.Assert(th.synced.Load(), qt.IsFalse)
}

func TestAppend(t *testing.T) {
	c := qt.New(t)
	th := new(TokenHolders).Init(MonkeysAddress, CONTRACT_TYPE_ERC20, MonkeysCreationBlock, 5, "")

	holderAddr := common.HexToAddress("0xe54d702f98E312aBA4318E3c6BDba98ab5e11012")
	holderBalance := new(big.Int).SetUint64(16000000000000000000)
	_, exists := th.holders.Load(holderAddr)
	c.Assert(exists, qt.IsFalse)
	th.Append(holderAddr, holderBalance)
	balance, exists := th.holders.Load(holderAddr)
	c.Assert(exists, qt.IsTrue)
	c.Assert(balance.(*big.Int).String(), qt.Equals, holderBalance.String())
}

func TestExists(t *testing.T) {
	c := qt.New(t)
	th := new(TokenHolders).Init(MonkeysAddress, CONTRACT_TYPE_ERC20, MonkeysCreationBlock, 5, "")

	holderAddr := common.HexToAddress("0xe54d702f98E312aBA4318E3c6BDba98ab5e11012")
	holderBalance := new(big.Int).SetUint64(16000000000000000000)
	c.Assert(th.Exists(holderAddr), qt.IsFalse)
	th.Append(holderAddr, holderBalance)
	c.Assert(th.Exists(holderAddr), qt.IsTrue)
}

func TestDel(t *testing.T) {
	c := qt.New(t)
	th := new(TokenHolders).Init(MonkeysAddress, CONTRACT_TYPE_ERC20, MonkeysCreationBlock, 5, "")

	holderAddr := common.HexToAddress("0xe54d702f98E312aBA4318E3c6BDba98ab5e11012")
	holderBalance := new(big.Int).SetUint64(16000000000000000000)
	th.Append(holderAddr, holderBalance)
	balance, exists := th.holders.Load(holderAddr)
	c.Assert(exists, qt.IsTrue)
	c.Assert(balance.(*big.Int).String(), qt.Equals, holderBalance.String())

	th.Del(holderAddr)
	notRemove, exists := th.holders.Load(holderAddr)
	c.Assert(exists, qt.IsTrue)
	c.Assert(notRemove.(bool), qt.IsFalse)
}

func TestFlushHolders(t *testing.T) {
	c := qt.New(t)
	th := new(TokenHolders).Init(MonkeysAddress, CONTRACT_TYPE_ERC20, MonkeysCreationBlock, 5, "")

	holderAddr := common.HexToAddress("0xe54d702f98E312aBA4318E3c6BDba98ab5e11012")
	holderBalance := new(big.Int).SetUint64(16000000000000000000)
	th.Append(holderAddr, holderBalance)
	balance, exists := th.holders.Load(holderAddr)
	c.Assert(exists, qt.IsTrue)
	c.Assert(balance.(*big.Int).String(), qt.Equals, holderBalance.String())

	th.FlushHolders()
	_, exists = th.holders.Load(holderAddr)
	c.Assert(exists, qt.IsFalse)
}

func TestBlockDone(t *testing.T) {
	c := qt.New(t)
	th := new(TokenHolders).Init(MonkeysAddress, CONTRACT_TYPE_ERC20, MonkeysCreationBlock, 5, "")

	_, exists := th.blocks.Load(MonkeysCreationBlock + 500)
	c.Assert(exists, qt.IsFalse)

	th.BlockDone(MonkeysCreationBlock + 500)
	processed, exists := th.blocks.Load(MonkeysCreationBlock + 500)
	c.Assert(exists, qt.IsTrue)
	c.Assert(processed.(bool), qt.IsTrue)
}

func TestHasBlock(t *testing.T) {
	c := qt.New(t)
	th := new(TokenHolders).Init(MonkeysAddress, CONTRACT_TYPE_ERC20, MonkeysCreationBlock, 5, "")

	c.Assert(th.HasBlock(MonkeysCreationBlock), qt.IsFalse)
	th.BlockDone(MonkeysCreationBlock)
	c.Assert(th.HasBlock(MonkeysCreationBlock), qt.IsTrue)
}

func TestLastBlock(t *testing.T) {
	c := qt.New(t)
	th := new(TokenHolders).Init(MonkeysAddress, CONTRACT_TYPE_ERC20, MonkeysCreationBlock, 5, "")

	c.Assert(th.LastBlock(), qt.Equals, MonkeysCreationBlock)
	th.BlockDone(MonkeysCreationBlock + 1)
	c.Assert(th.LastBlock(), qt.Equals, MonkeysCreationBlock+1)
	th.BlockDone(MonkeysCreationBlock + 2)
	c.Assert(th.LastBlock(), qt.Equals, MonkeysCreationBlock+2)
}

func TestSynced(t *testing.T) {
	c := qt.New(t)
	th := new(TokenHolders).Init(MonkeysAddress, CONTRACT_TYPE_ERC20, MonkeysCreationBlock, 5, "")

	c.Assert(th.synced.Load(), qt.IsFalse)
	th.Synced()
	c.Assert(th.synced.Load(), qt.IsTrue)
}

func TestIsSynced(t *testing.T) {
	c := qt.New(t)
	th := new(TokenHolders).Init(MonkeysAddress, CONTRACT_TYPE_ERC20, MonkeysCreationBlock, 5, "")

	c.Assert(th.IsSynced(), qt.IsFalse)
	th.Synced()
	c.Assert(th.IsSynced(), qt.IsTrue)
}
