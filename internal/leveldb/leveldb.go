package leveldb

import (
	"LevelGo/internal/config"
	"log"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type LevelDB struct {
	DB       *leveldb.DB
	lock     *sync.RWMutex
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
		DB:       db,
		lock:     new(sync.RWMutex),
		NotFound: leveldb.ErrNotFound,
	}
}

func (self *LevelDB) Get(key []byte) ([]byte, error) {
	self.lock.RLock()
	defer self.lock.RUnlock()
	data, err := self.DB.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (self *LevelDB) Set(key []byte, value []byte) error {
	self.lock.Lock()
	defer self.lock.Unlock()
	err := self.DB.Put(key, value, &opt.WriteOptions{Sync: true})
	return err
}

func (self *LevelDB) Has(key []byte) (bool, error) {
	self.lock.RLock()
	defer self.lock.RUnlock()
	has, err := self.DB.Has(key, nil)
	if err != nil {
		return false, err
	}
	return has, nil
}

func (self *LevelDB) Del(key []byte) error {
	self.lock.Lock()
	defer self.lock.Unlock()
	err := self.DB.Delete(key, &opt.WriteOptions{Sync: true})
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
	err := self.DB.Write(batch.batch, nil)
	return err
}

func (self *LevelDB) Close() {
	self.DB.Close()
}
