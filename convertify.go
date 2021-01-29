package convertify

import (
	"log"
	"reflect"
)

// Convert a map into a given item
func Convert(input map[string]interface{}, target interface{}) {
	if input == nil || len(input) == 0 || target == nil {
		return
	}

	t := reflect.TypeOf(target)
	fields := make([]reflect.StructField, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		fields[i] = t.Field(i)
	}
	log.Print(fields)
}
