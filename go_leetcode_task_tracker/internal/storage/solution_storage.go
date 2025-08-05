package storage

import "leetcode_task_tracker/internal/task"

type SolutionStorage interface {
	Init() error
	Initialized() bool
	AddSolution(t task.Task) error
	GetSolutions(id int) (task.Solutions, error)
	GetAllSolutions() map[int]task.Solutions
	UpdateSolution(t task.Task) error
	RemoveSolution(id int) error
}
