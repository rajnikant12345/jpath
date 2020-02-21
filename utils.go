package jpath

import (
	"reflect"
	"strconv"
)

func IsValidAddress(addr string) bool {
	return true
}

func isfloat(in interface{}) bool {
	return reflect.TypeOf(in).Kind() == reflect.Float64
}

func isstring(in interface{}) bool {
	return reflect.TypeOf(in).Kind() == reflect.String
}

func isslice(in interface{}) bool {
	return reflect.TypeOf(in).Kind() == reflect.Slice
}

func ismap(in interface{}) bool {
	return reflect.TypeOf(in).Kind() == reflect.Map
}

func isJsonPathSlice(in string) int {
	if len(in) >= 3 && in[0] == '[' && in[len(in)-1] == ']' {
		val := in[1 : len(in)-1]
		if val == "*" {
			return 0
		} else {
			val, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return -1
			}
			return int(val)
		}
	}
	return -1
}
