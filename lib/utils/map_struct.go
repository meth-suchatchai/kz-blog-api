package utils

import (
	"fmt"
	"reflect"
)

func FillMapStruct(input interface{}) map[string]string {
	values := reflect.ValueOf(input)
	types := values.Type()

	output := make(map[string]string)

	for i := 0; i < values.NumField(); i++ {
		output[types.Field(i).Name] = fmt.Sprintf("%v", values.Field(i))
	}

	return output
}
