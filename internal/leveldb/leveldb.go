package leveldb

import (
	"LevelGo/internal/config"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDB struct {
	db       *leveldb.DB
	NotFound error
}

type LevelDBBatchOperations struct {
	batch *leveldb.Batch
}

func GetLevelDB(leveldbConf *config.LevelDBSetting) *LevelDB {
	var err error
	db, err := leveldb.OpenFile(leveldbConf.File, nil)
	if err != nil {
		panic(err)
	}
	log.Print("LevelDB Init")
	return &LevelDB{
		db:       db,
		NotFound: leveldb.ErrNotFound,
	}
}

func (self *LevelDB) Get(key []byte) ([]byte, error) {
	data, err := self.db.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (self *LevelDB) Set(key []byte, value []byte) error {
	err := self.db.Put(key, value, nil)
	return err
}

func (self *LevelDB) Has(key []byte) (bool, error) {
	has, err := self.db.Has(key, nil)
	if err != nil {
		return false, err
	}
	return has, nil
}

func (self *LevelDB) Del(key []byte) error {
	err := self.db.Delete(key, nil)
	return err
}

func (self *LevelDB) GetBatch() *LevelDBBatchOperations {
	batch := new(leveldb.Batch)
	return &LevelDBBatchOperations{
		batch: batch,
	}
}

func (self *LevelDBBatchOperations) Set(key []byte, value []byte) {
	self.batch.Put(key, value)
}

func (self *LevelDBBatchOperations) Del(key []byte) {
	self.batch.Delete(key)
}

func (self *LevelDB) RunBatch(batch *LevelDBBatchOperations) error {
	err := self.db.Write(batch.batch, nil)
	return err
}

func (self *LevelDB) CloseLevelDB() {
	self.db.Close()
}
