package custom_errors

type TaskExistsError struct{}

func (e TaskExistsError) Error() string {
	return "task already exists"
}

type TaskNotFoundError struct{}

func (e TaskNotFoundError) Error() string {
	return "task is not found"
}
