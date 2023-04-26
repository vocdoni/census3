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

var (
	ErrCreatingCensusDB          = fmt.Errorf("error creating the census trees database")
	ErrInitializingIPFS          = fmt.Errorf("error initializing the IPFS service")
	ErrCreatingCensusTree        = fmt.Errorf("error creating the census tree")
	ErrSavingCensusTree          = fmt.Errorf("error saving the census tree")
	ErrAddingHoldersToCensusTree = fmt.Errorf("error adding holders to the census tree")
	ErrPublishingCensusTree      = fmt.Errorf("error publishing the census tree")
	ErrPruningCensusTree         = fmt.Errorf("error pruning the census tree")
)

type CensusDefinition struct {
	ID         int
	StrategyID int
	Type       models.Census_Type
	URI        string
	AuthToken  *uuid.UUID
	MaxLevels  int
	holders    map[common.Address]int
	tree       *censustree.Tree
}

func DefaultCensusDefinition(id, strategyID int, holders map[common.Address]int) *CensusDefinition {
	return &CensusDefinition{
		ID:         id,
		StrategyID: strategyID,
		Type:       defaultCensusType,
		URI:        "",
		AuthToken:  nil,
		MaxLevels:  defaultMaxLevels,
		holders:    holders,
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

type CensusDB struct {
	treeDB  db.Database
	storage storagelayer.Storage
}

func NewCensusDB(dataDir string) (*CensusDB, error) {
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
	return &CensusDB{treeDB: db, storage: storage}, nil
}

// create new census and add the holders
func (cdb *CensusDB) CreateAndPublish(def *CensusDefinition) (*CensusDump, error) {
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
	for addr, value := range def.holders {
		holdersAddresses = append(holdersAddresses, addr.Bytes())
		holdersValues = append(holdersValues, []byte(strconv.Itoa(value)))
	}
	// add the holders
	if _, err := def.tree.AddBatch(holdersAddresses, holdersValues); err != nil {
		log.Errorw(ErrAddingHoldersToCensusTree, err.Error())
		return nil, ErrAddingHoldersToCensusTree
	}
	// publish on IPFS
	dump, err := cdb.publish(def)
	if err != nil {
		log.Errorw(ErrPublishingCensusTree, err.Error())
		return nil, ErrPublishingCensusTree
	}
	// prune the created census from the database
	if err := cdb.delete(def); err != nil {
		log.Errorw(ErrPruningCensusTree, err.Error())
		return nil, ErrPruningCensusTree
	}
	return dump, nil
}

func (cdb *CensusDB) newTree(def *CensusDefinition) (*CensusDefinition, error) {
	tree, err := censustree.New(censustree.Options{
		Name:       censusDBKey(def.ID),
		ParentDB:   cdb.treeDB,
		MaxLevels:  def.MaxLevels,
		CensusType: def.Type})
	if err != nil {
		return nil, err
	}
	tree.Publish()
	def.tree = tree
	return def, nil
}

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

func (cdb *CensusDB) publish(def *CensusDefinition) (*CensusDump, error) {
	root, err := def.tree.Root()
	if err != nil {
		return nil, err
	}
	log.Infow("publishing census", "root", string(root))
	data, err := def.tree.Dump()
	if err != nil {
		return nil, err
	}

	dump := &CensusDump{
		ID:         def.ID,
		StrategyID: def.StrategyID,
		Type:       def.Type,
		RootHash:   root,
		Data:       compressor.NewCompressor().CompressBytes(data),
		MaxLevels:  def.MaxLevels,
	}
	exportData, err := json.Marshal(dump)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if dump.URI, err = cdb.storage.Publish(ctx, exportData); err != nil {
		return nil, err
	}
	return dump, nil
}

func (cdb *CensusDB) delete(def *CensusDefinition) error {
	wtx := cdb.treeDB.WriteTx()
	defer wtx.Discard()
	if err := wtx.Delete([]byte(censusDBKey(def.ID))); err != nil {
		return err
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
func censusDBKey(censusID int) string {
	return fmt.Sprintf("%s%x", censusDBprefix, []byte(strconv.Itoa(censusID)))
}

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
