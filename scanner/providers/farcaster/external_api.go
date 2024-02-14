package farcaster

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.vocdoni.io/dvote/log"
)

var (
	defaultAPICooldown    = time.Hour * 6
	verificationsByFidURI = "verificationsByFid?fid=%s"
	userDataByFidURI      = "userDataByFid?fid=%s&user_data_type=1"
)

type VerificationAddEthAddressBody struct {
	Address      string `json:"address"`
	EthSignature string `json:"ethSignature"`
	BlockHash    string `json:"blockHash"`
}

type MessageDataVerificationsByFID struct {
	Type                          string                        `json:"type"`
	Fid                           int                           `json:"fid"`
	Timestamp                     int64                         `json:"timestamp"`
	Network                       string                        `json:"network"`
	VerificationAddEthAddressBody VerificationAddEthAddressBody `json:"verificationAddEthAddressBody"`
}

type MessageVerificationsByFID struct {
	Data            MessageDataVerificationsByFID `json:"data"`
	Hash            string                        `json:"hash"`
	HashScheme      string                        `json:"hashScheme"`
	Signature       string                        `json:"signature"`
	SignatureScheme string                        `json:"signatureScheme"`
	Signer          string                        `json:"signer"`
}

// FarcasteAPIResponse is the response from the farcaster API call to verificationsByFid
type FarcasterAPIResponseVerificationsByFID struct {
	Messages      []MessageVerificationsByFID `json:"messages"`
	NextPageToken string                      `json:"nextPageToken"`
}

func (p *FarcasterProvider) apiVerificationsByFID(fid *big.Int, address common.Address, offset int,
) (*FarcasterAPIResponseVerificationsByFID, error) {
	// compose the endpoint for the request
	strURL, err := url.JoinPath(p.apiEndpoint, fmt.Sprintf(verificationsByFidURI, fid.String()))
	if err != nil {
		return nil, err
	}
	endpoint, err := url.Parse(strURL)
	if err != nil {
		return nil, err
	}
	q := endpoint.Query()
	// q.Add("limit", fmt.Sprint(FARCASTER_MAX_LIMIT))
	// q.Add("offset", fmt.Sprint(offset))
	endpoint.RawQuery = q.Encode()
	// create request and add headers
	req, err := http.NewRequest("GET", endpoint.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", p.accessToken)
	// do the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	// decode response
	rawResults, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Error(err)
		}
	}()
	response := FarcasterAPIResponseVerificationsByFID{}
	if err := json.Unmarshal(rawResults, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Define the struct for the "userDataBody"
type UserDataBodyUserDataByFID struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Define the struct for the "data" field
type DataUserDataByFID struct {
	Type         string                    `json:"type"`
	Fid          int                       `json:"fid"`
	Timestamp    int64                     `json:"timestamp"`
	Network      string                    `json:"network"`
	UserDataBody UserDataBodyUserDataByFID `json:"userDataBody"`
}

// Define the top-level struct for the entire response
type FarcasterAPIResponseUserDataByFID struct {
	Data            DataUserDataByFID `json:"data"`
	Hash            string            `json:"hash"`
	HashScheme      string            `json:"hashScheme"`
	Signature       string            `json:"signature"`
	SignatureScheme string            `json:"signatureScheme"`
	Signer          string            `json:"signer"`
}

func (p *FarcasterProvider) apiUserDataByFID(fid *big.Int) (*FarcasterAPIResponseUserDataByFID, error) {
	// compose the endpoint for the request
	strURL, err := url.JoinPath(p.apiEndpoint, fmt.Sprintf(userDataByFidURI, fid.String()))
	if err != nil {
		return nil, err
	}
	endpoint, err := url.Parse(strURL)
	if err != nil {
		return nil, err
	}
	q := endpoint.Query()
	// q.Add("limit", fmt.Sprint(FARCASTER_MAX_LIMIT))
	// q.Add("offset", fmt.Sprint(offset))
	endpoint.RawQuery = q.Encode()
	// create request and add headers
	req, err := http.NewRequest("GET", endpoint.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", p.accessToken)
	// do the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	// decode response
	rawResults, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Error(err)
		}
	}()
	response := FarcasterAPIResponseUserDataByFID{}
	if err := json.Unmarshal(rawResults, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
