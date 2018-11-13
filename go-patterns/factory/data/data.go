package data

import (
	"io"
	"os"
)

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	MemoryStorage
)

func NewStore(t StorageType) Store {
	switch t {
	case MemoryStorage:
		return newMemoryStorage()
	case DiskStorage:
		return newDiskStorage()
	}
	return nil
}

type memoryStorage struct {
}

func (p memoryStorage) Open(file string) (io.ReadWriteCloser, error) {
	f, err := os.Create(file)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func newMemoryStorage() Store {
	return memoryStorage{}
}

type diskStorage struct {
}

func newDiskStorage() Store {
	return diskStorage{}
}

func (p diskStorage) Open(file string) (io.ReadWriteCloser, error) {
	f, err := os.Create(file)
	if err != nil {
		return nil, err
	}
	return f, nil
}
