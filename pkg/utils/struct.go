package utils

import (
	"reflect"
)

func StructToKeyValue(in interface{}, tag string) ([]string, []interface{}) {
	var (
		keys   []string
		values []interface{}
	)

	rt := reflect.TypeOf(in)
	rv := reflect.ValueOf(in)

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	for i := 0; i < rv.NumField(); i++ {
		if !rv.Field(i).IsZero() {
			key := rt.Field(i).Tag.Get(tag)
			value := rv.Field(i)

			keys = append(keys, key)
			values = append(values, value.Interface())
		}
	}

	return keys, values
}
