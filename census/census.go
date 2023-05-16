package census

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math/big"
	"path/filepath"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"go.vocdoni.io/dvote/api/censusdb"
	"go.vocdoni.io/dvote/censustree"
	storagelayer "go.vocdoni.io/dvote/data"
	"go.vocdoni.io/dvote/data/ipfs"
	"go.vocdoni.io/dvote/data/ipfs/ipfsconnect"
	"go.vocdoni.io/dvote/db"
	"go.vocdoni.io/dvote/db/metadb"
	"go.vocdoni.io/dvote/log"
	"go.vocdoni.io/proto/build/go/models"
)

const (
	censusDBprefix    = "cs_"
	defaultMaxLevels  = censustree.DefaultMaxLevels
	defaultCensusType = models.Census_ARBO_BLAKE2B
)

var (
	ErrCreatingCensusDB          = fmt.Errorf("error creating the census trees database")
	ErrInitializingIPFS          = fmt.Errorf("error initializing the IPFS service")
	ErrCreatingCensusTree        = fmt.Errorf("error creating the census tree")
	ErrSavingCensusTree          = fmt.Errorf("error saving the census tree")
	ErrAddingHoldersToCensusTree = fmt.Errorf("error adding holders to the census tree")
	ErrPublishingCensusTree      = fmt.Errorf("error publishing the census tree")
	ErrPruningCensusTree         = fmt.Errorf("error pruning the census tree")
)

// CensusDefinintion envolves the required parameters to create and use a
// census merkle tree
type CensusDefinition struct {
	ID         int
	StrategyID int
	Type       models.Census_Type
	URI        string
	AuthToken  *uuid.UUID
	MaxLevels  int
	Holders    map[common.Address]*big.Int
	tree       *censustree.Tree
}

// DefaultCensusDefinition function returns a populated census definition with
// the default values for some parameters and the supplied values for the rest.
func DefaultCensusDefinition(id, strategyID int, holders map[common.Address]*big.Int) *CensusDefinition {
	return &CensusDefinition{
		ID:         id,
		StrategyID: strategyID,
		Type:       defaultCensusType,
		URI:        "",
		AuthToken:  nil,
		MaxLevels:  defaultMaxLevels,
		Holders:    holders,
	}
}

// CensusDump is a struct that contains the data of a census. It is used
// for import/export operations.
type CensusDump struct {
	// the following attributes are only for internal use and they have not be
	// published on IPFS
	ID         int    `json:"-"`
	StrategyID int    `json:"-"`
	URI        string `json:"-"`

	Type     models.Census_Type `json:"type"`
	RootHash []byte             `json:"rootHash"`
	Data     []byte             `json:"data"`
	// MaxLevels is required to load the census with the original size because
	// it could be different according to the election (and census) type.
	MaxLevels int `json:"maxLevels"`
}

type PublishedCensus struct {
	ID         int
	StrategyID int
	RootHash   []byte
	URI        string
	Dump       []byte
}

// CensusDB struct envolves the internal trees database and the IPFS handler,
// required to create and publish censuses.
type CensusDB struct {
	treeDB   db.Database
	storage  storagelayer.Storage
	ipfsConn *ipfsconnect.IPFSConnect
}

// NewCensusDB function instansiates an new internal tree database that will be
// located into the directory path provided.
func NewCensusDB(dataDir, groupKey string) (*CensusDB, error) {
	db, err := metadb.New(db.TypePebble, filepath.Join(dataDir, "censusdb"))
	if err != nil {
		log.Errorw(ErrCreatingCensusDB, err.Error())
		return nil, ErrCreatingCensusDB
	}
	ipfsConfig := storagelayer.IPFSNewConfig(dataDir)
	storage, err := storagelayer.Init(storagelayer.IPFS, ipfsConfig)
	if err != nil {
		log.Errorw(ErrInitializingIPFS, err.Error())
		return nil, ErrInitializingIPFS
	}
	var ipfsConn *ipfsconnect.IPFSConnect
	if len(groupKey) > 0 {
		ipfsConn = ipfsconnect.New(groupKey, storage.(*ipfs.Handler))
		ipfsConn.Start()
	}
	return &CensusDB{treeDB: db, storage: storage, ipfsConn: ipfsConn}, nil
}

// CreateAndPublish function creates a new census tree based on the definition
// provided and publishes it to IPFS. It needs to persist it temporaly into a
// internal trees database.
func (cdb *CensusDB) CreateAndPublish(def *CensusDefinition) (*PublishedCensus, error) {
	var err error
	if def, err = cdb.newTree(def); err != nil {
		log.Errorw(ErrCreatingCensusTree, err.Error())
		return nil, ErrCreatingCensusTree
	}
	// save the census definition into the trees database
	if err := cdb.save(def); err != nil {
		log.Errorw(ErrSavingCensusTree, err.Error())
		return nil, ErrSavingCensusTree
	}
	// encode the holders
	holdersAddresses, holdersValues := [][]byte{}, [][]byte{}
	for addr, balance := range def.Holders {
		value := def.tree.BigIntToBytes(balance)
		key, err := def.tree.Hash(addr.Bytes())
		if err != nil {
			return nil, ErrAddingHoldersToCensusTree
		}
		holdersAddresses = append(holdersAddresses, key[:censustree.DefaultMaxKeyLen])
		holdersValues = append(holdersValues, value)
	}
	// add the holders
	if _, err := def.tree.AddBatch(holdersAddresses, holdersValues); err != nil {
		log.Errorw(ErrAddingHoldersToCensusTree, err.Error())
		return nil, ErrAddingHoldersToCensusTree
	}
	// publish on IPFS
	res, err := cdb.publish(def)
	if err != nil {
		log.Errorw(ErrPublishingCensusTree, err.Error())
		return nil, ErrPublishingCensusTree
	}
	// prune the created census from the database
	if err := cdb.delete(def); err != nil {
		log.Errorw(ErrPruningCensusTree, err.Error())
		return nil, ErrPruningCensusTree
	}
	return res, nil
}

// newTree function creates a new census tree based on the provided definition
func (cdb *CensusDB) newTree(def *CensusDefinition) (*CensusDefinition, error) {
	var err error
	def.tree, err = censustree.New(censustree.Options{
		Name:       censusDBKey(def.ID),
		ParentDB:   cdb.treeDB,
		MaxLevels:  def.MaxLevels,
		CensusType: def.Type})
	if err != nil {
		return nil, err
	}
	def.tree.Publish()
	return def, nil
}

// save function persists the provided census definition into the internal trees
// database. It encodes the census definition using Gob and the creates a new
// entry on the database using the census ID as its identifier.
func (cdb *CensusDB) save(def *CensusDefinition) error {
	wtx := cdb.treeDB.WriteTx()
	defer wtx.Discard()
	defData := bytes.Buffer{}
	enc := gob.NewEncoder(&defData)
	if err := enc.Encode(def); err != nil {
		return err
	}
	if err := wtx.Set([]byte(censusDBKey(def.ID)), defData.Bytes()); err != nil {
		return err
	}
	return wtx.Commit()
}

// publish function takes a dump of the given census, serialises and publishes
// it to IPFS. If all goes well, it returns the census dump struct.
func (cdb *CensusDB) publish(def *CensusDefinition) (*PublishedCensus, error) {
	// get census tree root
	root, err := def.tree.Root()
	if err != nil {
		return nil, err
	}
	// get tree dump
	data, err := def.tree.Dump()
	if err != nil {
		return nil, err
	}
	// create census dump compressing the tree dump
	res := &PublishedCensus{
		ID:         def.ID,
		StrategyID: def.StrategyID,
		RootHash:   root,
	}
	// encode it into a JSON
	res.Dump, err = censusdb.BuildExportDump(root, data, def.Type, def.MaxLevels)
	if err != nil {
		return nil, err
	}
	// publish it on IPFS
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if res.URI, err = cdb.storage.Publish(ctx, res.Dump); err != nil {
		return nil, err
	}
	return res, nil
}

// delete function removes the census provided from the internal tree database
func (cdb *CensusDB) delete(def *CensusDefinition) error {
	wtx := cdb.treeDB.WriteTx()
	defer wtx.Discard()
	if err := wtx.Delete([]byte(censusDBKey(def.ID))); err != nil {
		return err
	}
	// the removal of the tree from the disk is done in a separate goroutine.
	// This is because the tree is locked and we don't want to block the
	// operations, and depending on the size of the tree, it can take a while
	// to delete it.
	go func() {
		_, err := censustree.DeleteCensusTreeFromDatabase(cdb.treeDB, censusDBKey(def.ID))
		if err != nil {
			log.Warnf("error deleting census tree %x: %s", def.ID, err)
		}
	}()
	return wtx.Commit()
}

// censusDBKey returns the db key of the census tree in the database given a censusID.
func censusDBKey(censusID int) string {
	return fmt.Sprintf("%s%x", censusDBprefix, []byte(strconv.Itoa(censusID)))
}

// TODO: Only used to debug on MVP stage, remove it
func (cdb *CensusDB) Check(def *CensusDefinition, root []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// download census dump from IPFS
	data, err := cdb.storage.Retrieve(ctx, def.URI, 0)
	if err != nil {
		return err
	}
	// decode result
	dump := CensusDump{}
	if err := json.Unmarshal(data, &dump); err != nil {
		return err
	}
	// compare roots
	if strDumpRoot := common.Bytes2Hex(dump.RootHash); strDumpRoot != string(root) {
		return fmt.Errorf("root hashes does not match (%s != %s)", string(root), strDumpRoot)
	}
	return nil
}
