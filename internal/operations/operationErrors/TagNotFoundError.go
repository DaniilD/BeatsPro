package operationErrors

import "BeatsPro/internal/Enums"

type TagNotFoundError struct {
	code int
}

func NewTagNotFoundError() *TagNotFoundError {
	error := &TagNotFoundError{}

	error.code = Enums.TAG_NOT_FOUND
	return error
}

func (error *TagNotFoundError) Error() string {
	return "Тэг не найден"
}
