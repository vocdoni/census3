package farcaster

import (
	"bytes"
	"encoding/gob"

	"github.com/ethereum/go-ethereum/common"
)

// Encode/Decode array
func serializeArray(data any) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func deserializeArray(serializedData []byte, arrayType []byte) (any, error) {
	var data any
	if bytes.Equal(arrayType, ArrayTypeAddress) {
		data = make([]common.Address, 0)
	} else {
		data = make([]common.Hash, 0)
	}
	buf := bytes.NewBuffer(serializedData)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
// given a FID returns the user data stored in the farcaster database
func (p *FarcasterProvider) getUserByFID(ctx context.Context, fid uint64) *FarcasterUserData {
	user, err := p.db.QueriesRO.GetUserByFID(ctx, fid)
	// if error is sql.ErrNoRows, the user does not exist print log.Warn
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Fatalf("Cannot read from DB %w", err)
		return nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		log.Warnf("User with FID %d not found", fid)
		return nil
	}
	fud := &FarcasterUserData{
		FID:             big.NewInt(int64(user.Fid)),
		Username:        user.Username,
		CustodyAddress:  common.BytesToAddress(user.CustodyAddress),
		RecoveryAddress: common.BytesToAddress(user.RecoveryAddress),
		Signer:          common.BytesToAddress(user.Signer),
	}
	// deserialize app keys since they are stored as bytes
	if len(user.AppKeys) != 0 {
		appKeys, eArray(user.AppKeys)
		if err != nil {
			log.Fatalf("Cannot deserialize app keys %w", err)
			return nil
		}
		fud.AppKeys = appKeys
	}
	return fud
}
*/
