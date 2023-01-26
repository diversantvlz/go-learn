package hw09structvalidator

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	panic("implement me")
}

func Validate(v interface{}) error {
	validationErrors := ValidationErrors{}

	refType := reflect.TypeOf(v)
	refValue := reflect.ValueOf(v)

	for i := 0; i < refType.NumField(); i++ {
		filed := refType.Field(i)

		conditions := strings.Split(filed.Tag.Get("validate"), "|")
		for _, condition := range conditions {
			conditionSlice := strings.Split(condition, ":")
			conditionName := conditionSlice[0]
			conditionValue, err := strconv.Atoi(conditionSlice[1])
			if err != nil {
				return err
			}
			switch conditionName {
			case "len":
				if refValue.Field(i).Len() != conditionValue {
					validationErrors = append(validationErrors, ValidationError{
						Err: errors.New("lenght validation failed"),
					})
				}
				break

			}
		}
	}

	return validationErrors
}
