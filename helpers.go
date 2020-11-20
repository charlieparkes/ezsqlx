package ezsqlx

import (
	"fmt"
	"reflect"
)

func in(item interface{}, arrayType interface{}) bool {
	arr := reflect.ValueOf(arrayType)
	if arr.Kind() != reflect.Array && arr.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Invalid data type: %v", arr.Kind()))
	}
	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

func filter(base []string, exclude []string) []string {
	filtered := []string{}
	for _, f := range base {
		if in(f, exclude) {
			continue
		}
		filtered = append(filtered, f)
	}
	return filtered
}


func wrapStrings(inputs []string, mark string) []string {
	outputs := []string{}
	for _, f := range inputs {
		outputs = append(outputs, mark+f+mark)
	}
	return outputs
}

func prependStrings(inputs []string, prefix string) []string {
	outputs := []string{}
	for _, f := range inputs {
		outputs = append(outputs, prefix+f)
	}
	return outputs
}