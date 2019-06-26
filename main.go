package otus_lesson4

import (
	"errors"
	"sort"
)

var (
	ErrSliceIsNil = errors.New("slice is nil")
	ErrSliceIsEmpty = errors.New("slice is empty")
	ErrLessFuncIsNil = errors.New("less function is nil")
)

func FindMax(slice []interface{}, less func (i, j int) bool) (interface{}, error) {
	if err := validateParams(slice, less); err != nil {
		return nil, err
	}

	sort.Slice(slice, less)
	return slice[0], nil
}

func validateParams(slice []interface{}, less func(i int, j int) bool) error {
	if slice == nil {
		return ErrSliceIsNil
	}

	if len(slice) == 0 {
		return ErrSliceIsEmpty
	}

	if less == nil {
		return ErrLessFuncIsNil
	}

	return nil
}
