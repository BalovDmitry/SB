package task

import "time"

type Task struct {
	UpdateTime  time.Time
	Completed   bool
	Id          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Solution    string `json:"solution" validate:"required"`
	Description string `json:"description" validate:"required"`
}
