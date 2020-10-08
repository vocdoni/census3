package entitybridge

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/vocdoni/tokenstate"
	"gitlab.com/vocdoni/go-dvote/crypto/ethereum"
	log "gitlab.com/vocdoni/go-dvote/log"
	"gitlab.com/vocdoni/go-dvote/types"
)

type EntityBridgeService struct {
	TokenState *tokenstate.Web3
	ENS        *ENS
	SignKey    *ethereum.SignKeys
	Gateway    string
}

func NewEntityBridgeService() *EntityBridgeService {
	return &EntityBridgeService{
		TokenState: new(tokenstate.Web3),
		ENS:        new(ENS),
	}
}
func (bs *EntityBridgeService) Init(ctx context.Context, web3Endpoint, gwEndpoint, contractAddress string, privKey string) error {
	// add sign key
	bs.SignKey = new(ethereum.SignKeys)
	bs.SignKey.AddHexKey(privKey)
	bs.Gateway = gwEndpoint
	// conect to eth and the contract
	if err := bs.TokenState.Init(ctx, web3Endpoint, contractAddress); err != nil {
		return err
	}
	return nil
}

func (bs *EntityBridgeService) CreateEntityMetadata() (string, error) {
	// get token data
	td, err := bs.TokenState.GetTokenData()
	if err != nil {
		return "", err
	}
	reqBody, err := json.Marshal(map[string]string{
		"method":  "addFile",
		"type":    "ipfs",
		"name":    "entity-metadata.json",
		"content": td.String(),
	})
	if err != nil {
		return "", err
	}

	// upload entity data to IPFS via gateway
	resp, err := http.Post(bs.Gateway,
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var metaResp *types.MetaResponse
	if err := json.Unmarshal(body, metaResp); err != nil {
		return "", nil
	}
	log.Infof("upload file uri: %s", metaResp.URI)

	// eid == token address
	eIDBytes, err := hex.DecodeString(td.Address)
	if err != nil {
		return "", err
	}
	var eIDBytes32 [32]byte
	copy(eIDBytes32[:], eIDBytes)

	// set ens record
	if err := bs.ENS.SetText(bs.SignKey, eIDBytes32, "vnd.vocdoni.meta", metaResp.URI); err != nil {
		log.Warnf("cannot set entity metadata: %s", err)
		return "", err
	}
	log.Info("entity created successfully")
	return metaResp.URI, nil
}
