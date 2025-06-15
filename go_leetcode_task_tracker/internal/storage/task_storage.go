package storage

import "leetcode_task_tracker/internal/task"

type TaskStorage interface {
	Init()
	AddTask(task task.Task) error
	GetTask(id int) (task.Task, error)
	GetAllTasks() map[int]task.Task
	UpdateTask(task task.Task) error
	RemoveTask(id int) error
}
