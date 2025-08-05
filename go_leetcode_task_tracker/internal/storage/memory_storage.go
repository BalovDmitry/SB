package storage

import (
	"leetcode_task_tracker/internal/custom_errors"
	"leetcode_task_tracker/internal/task"
	"time"
)

type MemoryStorage struct {
	solutions   map[int]task.Solutions
	initialized bool
}

func (s *MemoryStorage) Init() error {
	s.solutions = make(map[int]task.Solutions)
	s.initialized = true
	return nil
}

func (s *MemoryStorage) Initialized() bool {
	return s.initialized
}

func (s *MemoryStorage) AddSolution(t task.Task) error {
	if _, found := s.solutions[t.Id]; found {
		return custom_errors.TaskExistsError{}
	}

	solutions := task.Solutions{}
	solutions.Id = t.Id
	solutions.Description = t.Description
	solutions.Name = t.Name

	solutions.LanguageToSolution = make(map[string]task.Solution)
	solutions.LanguageToSolution[t.Language] =
		task.Solution{UpdateTime: time.Now().Format(time.DateTime), Value: t.Solution}

	s.solutions[t.Id] = solutions
	return nil
}

func (s *MemoryStorage) GetSolutions(id int) (task.Solutions, error) {
	if _, found := s.solutions[id]; !found {
		return task.Solutions{}, custom_errors.TaskNotFoundError{}
	}

	return s.solutions[id], nil
}

func (s *MemoryStorage) GetAllSolutions() map[int]task.Solutions {
	return s.solutions
}

func (s *MemoryStorage) UpdateSolution(t task.Task) error {
	if _, found := s.solutions[t.Id]; !found {
		return custom_errors.TaskNotFoundError{}
	}

	solutionToUpdate := s.solutions[t.Id]
	if val, found := solutionToUpdate.LanguageToSolution[t.Language]; found {
		val.UpdateTime = time.Now().Format(time.DateTime)
		val.Value = t.Solution
		solutionToUpdate.LanguageToSolution[t.Language] = val
	} else {
		solutionToUpdate.LanguageToSolution[t.Language] =
			task.Solution{UpdateTime: time.Now().Format(time.DateTime), Value: t.Solution}
	}

	s.solutions[t.Id] = solutionToUpdate
	return nil
}

func (s *MemoryStorage) RemoveSolution(id int) error {
	if _, found := s.solutions[id]; !found {
		return custom_errors.TaskNotFoundError{}
	}

	delete(s.solutions, id)
	return nil
}
