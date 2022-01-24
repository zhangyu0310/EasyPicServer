package leveldb

import (
	"easyPicServer/store"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type LevelDB struct {
	DB *leveldb.DB
}

type Iterator struct {
	Iter *iterator.Iterator
}

func (db *LevelDB) Init(options ...interface{}) (err error) {
	db.DB, err = leveldb.OpenFile(options[0].(string),
		(options[1]).(*opt.Options))
	return
}

func (db *LevelDB) Get(key []byte, options ...interface{}) (value []byte, err error) {
	value, err = db.DB.Get(key, options[0].(*opt.ReadOptions))
	return
}

func (db *LevelDB) Set(key []byte, value []byte, options ...interface{}) (err error) {
	err = db.DB.Put(key, value, options[0].(*opt.WriteOptions))
	return
}

func (db *LevelDB) Delete(key []byte, options ...interface{}) (err error) {
	err = db.DB.Delete(key, options[0].(*opt.WriteOptions))
	return
}

func (db *LevelDB) Iterator() store.Iterator {
	it := db.DB.NewIterator(nil, nil)
	return &Iterator{Iter: &it}
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
