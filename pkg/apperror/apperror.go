package apperror

type AppError struct {
	Code    int
	Message string
	Detail  interface{}
}

func (e *AppError) Error() string {
	return e.Message
}

func NotFound(msg string) *AppError {
	return &AppError{Code: 404, Message: msg}
}

func BadRequest(msg string) *AppError {
	return &AppError{Code: 400, Message: msg}
}

func Unauthorized(msg string) *AppError {
	return &AppError{Code: 401, Message: msg}
}

func Forbidden(msg string) *AppError {
	return &AppError{Code: 403, Message: msg}
}

func Internal(msg string) *AppError {
	return &AppError{Code: 500, Message: msg}
}