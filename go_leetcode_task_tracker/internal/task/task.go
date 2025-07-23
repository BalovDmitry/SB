package task

type Task struct {
	UpdateTime  string `json:"time"`
	Completed   bool
	Id          int    `json:"id" validate:"required,min=0"`
	Name        string `json:"name" validate:"required"`
	Solution    string `json:"solution" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type View struct {
	IsSingle   bool
	Tasks      []Task
	TaskSingle Task
}
