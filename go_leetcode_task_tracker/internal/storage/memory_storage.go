package storage

import (
	"leetcode_task_tracker/internal/custom_errors"
	"leetcode_task_tracker/internal/task"
	"time"
)

type MemoryStorage struct {
	tasks map[int]task.Task
}

func (s *MemoryStorage) Init() error {
	s.tasks = make(map[int]task.Task)
	return nil
}

func (s *MemoryStorage) AddTask(task task.Task) error {
	if _, found := s.tasks[task.Id]; found {
		return custom_errors.TaskExistsError{}
	}

	task.Completed = true
	task.UpdateTime = time.Now().Format(time.DateTime)
	s.tasks[task.Id] = task
	return nil
}

func (s *MemoryStorage) GetTask(id int) (task.Task, error) {
	if _, found := s.tasks[id]; !found {
		return task.Task{}, custom_errors.TaskNotFoundError{}
	}

	return s.tasks[id], nil
}

func (s *MemoryStorage) GetAllTasks() map[int]task.Task {
	return s.tasks
}

func (s *MemoryStorage) UpdateTask(task task.Task) error {
	if _, found := s.tasks[task.Id]; !found {
		return custom_errors.TaskNotFoundError{}
	}

	taskToUpdate := s.tasks[task.Id]
	task.UpdateTime = time.Now().Format(time.DateTime)
	taskToUpdate.Solution = task.Solution

	s.tasks[task.Id] = taskToUpdate
	return nil
}

func (s *MemoryStorage) RemoveTask(id int) error {
	if _, found := s.tasks[id]; !found {
		return custom_errors.TaskNotFoundError{}
	}

	delete(s.tasks, id)
	return nil
}
