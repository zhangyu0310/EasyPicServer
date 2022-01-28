package redis

import (
	"easyPicServer/store"
	"errors"
	"github.com/go-redis/redis"
)

type Redis struct {
	rdb *redis.Client
}

type Iterator struct {
	rdb   *redis.Client
	iter  []string
	index int
}

func (rdb *Redis) Init(options ...interface{}) error {
	if len(options) < 1 {
		return errors.New("wrong number of arguments for initialize")
	} else {
		rdb.rdb = redis.NewClient(options[0].(*redis.Options))
	}
	return nil
}

func (rdb *Redis) Get(key []byte, _ ...interface{}) ([]byte, error) {
	value, err := rdb.rdb.Get(string(key)).Result()
	return []byte(value), err
}

func (rdb *Redis) Set(key []byte, value []byte, _ ...interface{}) error {
	return rdb.rdb.Set(string(key), value, 0).Err()
}

func (rdb *Redis) Delete(key []byte, _ ...interface{}) error {
	return rdb.rdb.Del(string(key)).Err()
}

func (rdb *Redis) Iterator(options ...interface{}) store.Iterator {
	var pattern string
	if len(options) >= 1 && options[0] != nil {
		pattern = options[0].(string)
	} else {
		pattern = "*"
	}
	iter, err := rdb.rdb.Keys(pattern).Result()
	if err != nil {
		iter = nil
	}
	return &Iterator{rdb: rdb.rdb, iter: iter, index: 0}
}

func (rdb *Redis) Close(...interface{}) error {
	return rdb.rdb.Close()
}

func (it *Iterator) Next() bool {
	if !it.Valid() {
		return false
	}
	if len(it.iter) > it.index {
		it.index++
		return true
	}
	return false
}

func (it *Iterator) Prev() bool {
	if !it.Valid() {
		return false
	}
	if it.index > 0 {
		it.index--
		return true
	}
	return false
}

func (it *Iterator) Seek(key []byte) bool {
	if !it.Valid() {
		return false
	}
	for i, v := range it.iter {
		if v == string(key) {
			it.index = i
			return true
		}
	}
	return false
}

func (it *Iterator) First() bool {
	if !it.Valid() {
		return false
	}
	it.index = 0
	return true
}

func (it *Iterator) Last() bool {
	if !it.Valid() {
		return false
	}
	if len(it.iter) > 0 {
		it.index = 0
	} else {
		it.index = len(it.iter) - 1
	}
	return true
}

func (it *Iterator) Key() []byte {
	if !it.Valid() {
		return []byte("")
	}
	return []byte(it.iter[it.index])
}

func (it *Iterator) Value() []byte {
	if !it.Valid() {
		return []byte("")
	}
	value, err := it.rdb.Get(string(it.Key())).Result()
	if err != nil {
		return []byte("")
	}
	return []byte(value)
}

func (it *Iterator) Valid() bool {
	return it.iter != nil
}

func (it *Iterator) Release() {
	return
}
