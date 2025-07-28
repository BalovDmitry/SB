package storage

import "leetcode_task_tracker/internal/task"

type SolutionStorage interface {
	Init() error
	AddTask(task task.Task) error
	GetTask(id int) (task.Task, error)
	GetAllTasks() map[int]task.Task
	UpdateTask(task task.Task) error
	RemoveTask(id int) error
}
