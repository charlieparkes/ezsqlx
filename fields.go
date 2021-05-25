package ezsqlx

import (
	"reflect"
)

func Columns(sf []reflect.StructField) []string {
	fields := []string{}
	for _, f := range sf {
		if tag := f.Tag.Get("db"); tag == "-" {
			continue
		} else if tag == "" {
			fields = append(fields, snakeCase(f.Name))
		} else {
			fields = append(fields, tag)
		}
	}
	return fields
}

func primaryKey(sf []reflect.StructField) []string {
	fields := []string{}
	for _, f := range sf {
		if f.Tag.Get("constraint") == "pk" {
			name := f.Tag.Get("db")
			if name == "" {
				name = snakeCase(f.Name)
			}
			fields = append(fields, name)
		}
	}
	return fields
}
