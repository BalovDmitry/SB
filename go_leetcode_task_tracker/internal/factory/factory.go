package factory

import (
	"leetcode_task_tracker/internal/handler"
	"leetcode_task_tracker/internal/storage"
)

type StorageType int

const (
	MemoryStorage StorageType = iota
)

func CreateStorage(storageType StorageType) storage.TaskStorage {
	switch storageType {
	case MemoryStorage:
		return &storage.MemoryTaskStorage{}
	default:
		return nil
	}
}

func CreateHandler(s storage.TaskStorage) handler.BaseHandler {
	switch s.(type) {
	case *storage.MemoryTaskStorage:
		return &handler.TaskStorageHandler{}
	default:
		return nil
	}
}
