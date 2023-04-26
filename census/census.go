package census

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"go.vocdoni.io/dvote/censustree"
	storagelayer "go.vocdoni.io/dvote/data"
	"go.vocdoni.io/dvote/data/compressor"
	"go.vocdoni.io/dvote/db"
	"go.vocdoni.io/dvote/db/metadb"
	"go.vocdoni.io/dvote/log"
	"go.vocdoni.io/proto/build/go/models"
)

const (
	censusDBprefix          = "cs_"
	censusDBreferencePrefix = "cr_"
	defaultMaxLevels        = 160
	defaultCensusType       = models.Census_ARBO_BLAKE2B
)

type CensusDefinition struct {
	ID        []byte
	Type      models.Census_Type
	URI       string
	AuthToken *uuid.UUID
	MaxLevels int
	holders   map[common.Address]int
}

func DefaultCensusDefinition(id int, holders map[common.Address]int) *CensusDefinition {
	return &CensusDefinition{
		ID:        []byte(strconv.Itoa(id)),
		Type:      defaultCensusType,
		URI:       "",
		AuthToken: nil,
		MaxLevels: defaultMaxLevels,
		holders:   holders,
	}
}

type CensusDB struct {
	treeDB  db.Database
	storage storagelayer.Storage
}

type Census struct {
	ID         int
	StrategyID int
	Root       []byte
	URI        string
}

// CensusDump is a struct that contains the data of a census. It is used
// for import/export operations.
type CensusDump struct {
	Type     models.Census_Type `json:"type"`
	RootHash []byte             `json:"rootHash"`
	Data     []byte             `json:"data"`
	// MaxLevels is required to load the census with the original size because
	// it could be different according to the election (and census) type.
	MaxLevels int `json:"maxLevels"`
}

func NewCensusDB(dataDir string) (*CensusDB, error) {
	db, err := metadb.New(db.TypePebble, filepath.Join(dataDir, "censusdb"))
	if err != nil {
		return nil, err
	}
	ipfsConfig := storagelayer.IPFSNewConfig(dataDir)
	storage, err := storagelayer.Init(storagelayer.IPFS, ipfsConfig)
	if err != nil {
		return nil, fmt.Errorf("error initializing IPFS handler: %w", err)
	}

	return &CensusDB{treeDB: db, storage: storage}, nil
}

// create new census and add the holders
func (cdb *CensusDB) CreateAndPublish(def *CensusDefinition) (*Census, error) {
	tree, err := cdb.initCensusTree(def)
	if err != nil {
		return nil, err
	}
	// save the census definition into the trees database
	if err := cdb.saveCensusDef(def); err != nil {
		return nil, err
	}
	// encode the holders
	holdersAddresses, holdersValues := [][]byte{}, [][]byte{}
	for addr, value := range def.holders {
		holdersAddresses = append(holdersAddresses, addr.Bytes())
		holdersValues = append(holdersValues, []byte(strconv.Itoa(value)))
	}
	// add the holders
	if _, err := tree.AddBatch(holdersAddresses, holdersValues); err != nil {
		return nil, fmt.Errorf("error including holders on censustree: %w", err)
	}
	// publish on IPFS
	uri, err := cdb.publishDump(def, tree)
	if err != nil {
		return nil, err
	}
	intID, err := strconv.Atoi(string(def.ID))
	if err != nil {
		return nil, err
	}
	root, err := tree.Root()
	if err != nil {
		return nil, err
	}
	res := &Census{
		ID:   intID,
		Root: root,
		URI:  "ipfs://" + uri,
	}
	// prune the created census from the database
	if err := cdb.deleteCensusDef(def); err != nil {
		return nil, err
	}
	return res, nil
}

func (cdb *CensusDB) initCensusTree(def *CensusDefinition) (*censustree.Tree, error) {
	tree, err := censustree.New(censustree.Options{
		Name:       censusDBKey(def.ID),
		ParentDB:   cdb.treeDB,
		MaxLevels:  def.MaxLevels,
		CensusType: def.Type})
	if err != nil {
		return nil, fmt.Errorf("errror creating new censustree: %w", err)
	}
	tree.Publish()
	return tree, nil
}

func (cdb *CensusDB) saveCensusDef(def *CensusDefinition) error {
	wtx := cdb.treeDB.WriteTx()
	defer wtx.Discard()
	defData := bytes.Buffer{}
	enc := gob.NewEncoder(&defData)
	if err := enc.Encode(def); err != nil {
		return err
	}
	if err := wtx.Set(append([]byte(censusDBreferencePrefix), def.ID...),
		defData.Bytes()); err != nil {
		return fmt.Errorf("error saving census definition: %w", err)
	}
	return wtx.Commit()
}

func (cdb *CensusDB) publishDump(def *CensusDefinition, tree *censustree.Tree) (string, error) {
	root, err := tree.Root()
	if err != nil {
		return "", err
	}
	log.Infow("publishing census", "root", string(root))
	data, err := tree.Dump()
	if err != nil {
		return "", err
	}

	dump := CensusDump{
		Type:      def.Type,
		RootHash:  root,
		Data:      compressor.NewCompressor().CompressBytes(data),
		MaxLevels: def.MaxLevels,
	}
	exportData, err := json.Marshal(dump)
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return cdb.storage.Publish(ctx, exportData)
}

func (cdb *CensusDB) deleteCensusDef(def *CensusDefinition) error {
	wtx := cdb.treeDB.WriteTx()
	defer wtx.Discard()
	if err := wtx.Delete(append([]byte(censusDBreferencePrefix), def.ID...)); err != nil {
		return fmt.Errorf("error pruning the census database: %w", err)
	}
	// the removal of the tree from the disk is done in a separate goroutine.
	// This is because the tree is locked and we don't want to block the operations,
	// and depending on the size of the tree, it can take a while to delete it.
	go func() {
		_, err := censustree.DeleteCensusTreeFromDatabase(cdb.treeDB, censusDBKey(def.ID))
		if err != nil {
			log.Warnf("error deleting census tree %x: %s", def.ID, err)
		}
	}()
	return wtx.Commit()
}

// censusDBKey returns the db key of the census tree in the database given a censusID.
func censusDBKey(censusID []byte) string {
	return fmt.Sprintf("%s%x", censusDBprefix, censusID)
}

func (cdb *CensusDB) Check(def *CensusDefinition, root []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	log.Infow("checking census", "id", string(def.ID), "root", string(root), "uri", def.URI)
	// download census dump from IPFS
	data, err := cdb.storage.Retrieve(ctx, def.URI, 0)
	if err != nil {
		return err
	}
	dump := CensusDump{}
	if err := json.Unmarshal(data, &dump); err != nil {
		return nil
	}

	if strDumpRoot := common.Bytes2Hex(dump.RootHash); strDumpRoot != string(root) {
		log.Warnw("root hashes does not match", "provided", string(root), "downloaded", strDumpRoot)
		return fmt.Errorf("root hashes does not match")
	}
	return nil
}
