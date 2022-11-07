package exceptions

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var Unauthorized = errors.New("unauthorized")
var EntityNotFound = errors.New("not found")
var Invalid = errors.New("invalid")

var InvalidAuthToken = fmt.Errorf("%w authentication token", Invalid)
var InvalidCredentials = fmt.Errorf("%w credentials", Invalid)

func GenericNotFound(model any) error {
	modelType := reflect.TypeOf(model).String()
	modelName := strings.ToLower(strings.Split(modelType, ".")[1])
	return fmt.Errorf("%s %w", modelName, EntityNotFound)
}
