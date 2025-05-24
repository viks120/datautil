package datautil

import (
	"errors"
)

func InsertSliceAt[T any](original, toInsert []T, position int) ([]T, error) {
	if position < 1 || position > len(original)+1 {
		return nil, errors.New("position out of range")
	}

	result := make([]T, 0, len(original)+len(toInsert))

	result = append(result, original[:position-1]...)
	result = append(result, toInsert...)
	if position-1 < len(original) {
		result = append(result, original[position-1:]...)
	}

	return result, nil
}
