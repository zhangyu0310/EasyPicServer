package leveldb

import (
	"easyPicServer/store"
	"errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type LevelDB struct {
	db *leveldb.DB
}

type Iterator struct {
	Iter *iterator.Iterator
}

func (db *LevelDB) Init(options ...interface{}) (err error) {
	if len(options) <= 0 {
		err = errors.New("wrong number of arguments for initialize")
	} else if len(options) <= 1 || (len(options) > 1 && options[1] == nil) {
		db.db, err = leveldb.OpenFile(options[0].(string), nil)
	} else {
		db.db, err = leveldb.OpenFile(options[0].(string), options[1].(*opt.Options))
	}
	return
}

func (db *LevelDB) Get(key []byte, options ...interface{}) (value []byte, err error) {
	if len(options) > 0 && options[0] != nil {
		value, err = db.db.Get(key, options[0].(*opt.ReadOptions))
	} else {
		value, err = db.db.Get(key, nil)
	}
	return
}

func (db *LevelDB) Set(key []byte, value []byte, options ...interface{}) (err error) {
	if len(options) > 0 && options[0] != nil {
		err = db.db.Put(key, value, options[0].(*opt.WriteOptions))
	} else {
		err = db.db.Put(key, value, nil)
	}
	return
}

func (db *LevelDB) Delete(key []byte, options ...interface{}) (err error) {
	if len(options) > 0 && options[0] != nil {
		err = db.db.Delete(key, options[0].(*opt.WriteOptions))
	} else {
		err = db.db.Delete(key, nil)
	}
	return
}

func (db *LevelDB) Iterator(options ...interface{}) store.Iterator {
	var slice *util.Range
	var ro *opt.ReadOptions
	if len(options) > 0 && options[0] != nil {
		slice = options[0].(*util.Range)
	}
	if len(options) > 1 && options[1] != nil {
		ro = options[1].(*opt.ReadOptions)
	}
	it := db.db.NewIterator(slice, ro)
	return &Iterator{Iter: &it}
}

func (db *LevelDB) Close(...interface{}) error {
	return db.db.Close()
}

func (it *Iterator) Next() bool {
	return (*it.Iter).Next()
}

func (it *Iterator) Prev() bool {
	return (*it.Iter).Prev()
}

func (it *Iterator) Seek(key []byte) bool {
	return (*it.Iter).Seek(key)
}

func (it *Iterator) First() bool {
	return (*it.Iter).First()
}

func (it *Iterator) Last() bool {
	return (*it.Iter).Last()
}

func (it *Iterator) Key() []byte {
	return (*it.Iter).Key()
}

func (it *Iterator) Value() []byte {
	return (*it.Iter).Value()
}

func (it *Iterator) Valid() bool {
	return (*it.Iter).Valid()
}

func (it *Iterator) Release() {
	(*it.Iter).Release()
}
