package custom_errors

type TaskExistsError struct{}

func (e TaskExistsError) Error() string {
	return "task already exists"
}

type TaskNotFoundError struct{}

func (e TaskNotFoundError) Error() string {
	return "task is not found"
}

type InitHandlerError struct{}

func (e InitHandlerError) Error() string {
	return "failed to initialize handler"
}

type RequestParserError struct{}

func (e RequestParserError) Error() string {
	return "can't receive data from request"
}
