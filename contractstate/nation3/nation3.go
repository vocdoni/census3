package nation3

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	nation3Passportcontracts "github.com/vocdoni/tokenstate/contracts/nation3/passport"
	nation3Tokencontracts "github.com/vocdoni/tokenstate/contracts/nation3/token"
	nation3VestedTokencontracts "github.com/vocdoni/tokenstate/contracts/nation3/vestedToken"
	"github.com/vocdoni/tokenstate/contractstate"
	"go.vocdoni.io/dvote/log"
	"go.vocdoni.io/dvote/util"
)

const (
	PASSPORT = iota
	VENATION
	NATION3
)

type Contracts struct {
	Passport *nation3Passportcontracts.Nation3PassportcontractsCaller
	VeNation *nation3VestedTokencontracts.Nation3VestedTokencontractsCaller
	Nation3  *nation3Tokencontracts.Nation3TokencontractsCaller
	PassportAddress,
	VeNationAddress,
	Nation3Address common.Address
}

type Nation3 struct {
	client    *ethclient.Client
	contacts  *Contracts
	networkID *big.Int
	close     chan (bool)
}

// Init creates and client connection and connects to all Nation3 contracts
// First contract is the Passport contract, second is the VeNation contract, third is the Nation3 contract
// Please respect this order when instantiating the contract addresses
func (n *Nation3) Init(ctx context.Context, web3Endpoint string, contractAddresses [3]string) error {
	var err error
	// connect to ethereum endpoint
	n.client, err = ethclient.Dial(web3Endpoint)
	if err != nil {
		log.Fatal(err)
	}
	n.networkID, err = n.client.ChainID(ctx)
	if err != nil {
		return err
	}
	log.Debugf("found ethereum network id %s", n.networkID.String())

	n.contacts = &Contracts{}
	// passport contract
	c, err := hex.DecodeString(util.TrimHex(contractAddresses[PASSPORT]))
	if err != nil {
		return err
	}
	caddr := common.Address{}
	caddr.SetBytes(c)
	if n.contacts.Passport, err = nation3Passportcontracts.NewNation3PassportcontractsCaller(caddr, n.client); err != nil {
		return err
	}
	n.contacts.PassportAddress = caddr
	log.Infof("loaded passport contract %s", caddr.String())

	// veNation token contract
	c, err = hex.DecodeString(util.TrimHex(contractAddresses[VENATION]))
	if err != nil {
		return err
	}
	caddr = common.Address{}
	caddr.SetBytes(c)
	if n.contacts.VeNation, err = nation3VestedTokencontracts.NewNation3VestedTokencontractsCaller(caddr, n.client); err != nil {
		return err
	}
	n.contacts.VeNationAddress = caddr
	log.Infof("loaded veNation contract %s", caddr.String())

	// nation3 token contract
	c, err = hex.DecodeString(util.TrimHex(contractAddresses[NATION3]))
	if err != nil {
		return err
	}
	caddr = common.Address{}
	caddr.SetBytes(c)
	if n.contacts.Nation3, err = nation3Tokencontracts.NewNation3TokencontractsCaller(caddr, n.client); err != nil {
		return err
	}
	n.contacts.Nation3Address = caddr
	log.Infof("loaded nation3 token contract %s", caddr.String())
	return nil
}

// Close closes the Nation3 client connection
func (n *Nation3) Close() {
	n.close <- true
}

// GetTokenData returns the token data for the given token operation
func (n *Nation3) GetTokenData(op uint8) (*contractstate.TokenData, error) {
	var err error
	td := &contractstate.TokenData{}

	switch op {
	case PASSPORT:
		td.Address = n.contacts.PassportAddress.Hex()
		td.Name, err = n.contacts.Passport.Name(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token name: %s", err)
		}
		td.Symbol, err = n.contacts.Passport.Symbol(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token symbol: %s", err)
		}
		td.TotalSupply, err = n.contacts.Passport.TotalSupply(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token total supply: %s", err)
		}
	case VENATION:
		td.Address = n.contacts.VeNationAddress.Hex()
		td.Name, err = n.contacts.VeNation.Name(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token name: %s", err)
		}
		td.Symbol, err = n.contacts.VeNation.Symbol(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token symbol: %s", err)
		}
		decimalsBig, err := n.contacts.VeNation.Decimals(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token decimals: %s", err)
		}
		td.Decimals = uint8(decimalsBig.Uint64())
		td.TotalSupply, err = n.contacts.VeNation.TotalSupply(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token total supply: %s", err)
		}
	case NATION3:
		td.Address = n.contacts.Nation3Address.Hex()
		td.Name, err = n.contacts.Nation3.Name(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token name: %s", err)
		}
		td.Symbol, err = n.contacts.Nation3.Symbol(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token symbol: %s", err)
		}
		td.Decimals, err = n.contacts.Nation3.Decimals(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token decimals: %s", err)
		}
		td.TotalSupply, err = n.contacts.Nation3.TotalSupply(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token total supply: %s", err)
		}

	default:
		return nil, fmt.Errorf("invalid operation")
	}

	return td, nil
}

// BalanceOfOrAt returns the balance of the given address, if op = VENATION you can get the
// balance at a the given block
func (n *Nation3) BalanceOfOrAt(ctx context.Context, op uint8, address string, atBlock *big.Int) (*big.Int, error) {
	var err error
	var balance *big.Int

	switch op {
	case PASSPORT:
		balance, err = n.contacts.Passport.BalanceOf(nil, common.HexToAddress(address))
		if err != nil {
			return nil, fmt.Errorf("unable to get passport balance: %s", err)
		}
	case VENATION:
		balance, err = n.contacts.VeNation.BalanceOfAt(nil, common.HexToAddress(address), atBlock)
		if err != nil {
			return nil, fmt.Errorf("unable to get veNation balance: %s", err)
		}
	case NATION3:
		balance, err = n.contacts.Nation3.BalanceOf(nil, common.HexToAddress(address))
		if err != nil {
			return nil, fmt.Errorf("unable to get nation3 balance: %s", err)
		}
	default:
		return nil, fmt.Errorf("invalid operation")
	}

	return balance, nil
}
