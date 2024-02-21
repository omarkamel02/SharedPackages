package SharedPackages

type ApplicationResult[T any] struct {
	Data       T
	StatusCode int
	Message    string
}

func NewApplicationResult[T any](data T, statusCode int, message string) ApplicationResult[T] {
	return ApplicationResult[T]{Data: data, StatusCode: statusCode, Message: message}
}
