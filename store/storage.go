package store

type Storage interface {
	Init(...interface{}) error
	Get([]byte, ...interface{}) ([]byte, error)
	Set([]byte, []byte, ...interface{}) error
	Delete([]byte, ...interface{}) error
	Iterator(options ...interface{}) Iterator
	Close(...interface{}) error
}

type Iterator interface {
	Next() bool
	Prev() bool
	Seek(key []byte) bool
	First() bool
	Last() bool
	Key() []byte
	Value() []byte
	Valid() bool
	Release()
}

var storage *Storage

func InitializeStorage(s Storage) {
	storage = &s
}

func GetStorage() *Storage {
	return storage
}
