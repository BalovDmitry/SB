package factory

import (
	"leetcode_task_tracker/internal/handler"
	"leetcode_task_tracker/internal/storage"
)

type StorageType int

const (
	MemoryStorage StorageType = iota
)

func CreateStorage(storageType StorageType) storage.SolutionStorage {
	switch storageType {
	case MemoryStorage:
		return &storage.MemoryStorage{}
	default:
		return nil
	}
}

func CreateHandler(s storage.SolutionStorage) handler.BaseHandler {
	switch s.(type) {
	case *storage.MemoryStorage:
		return &handler.SolutionStorageHandler{}
	default:
		return nil
	}
}
