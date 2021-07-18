package operationErrors

import "BeatsPro/internal/Enums"

type TagNotFoundError struct {
	Code int
}

func NewTagNotFoundError() *TagNotFoundError {
	error := &TagNotFoundError{}

	error.Code = Enums.TAG_NOT_FOUND
	return error
}

func (error *TagNotFoundError) Error() string {
	return "Тэг не найден"
}
