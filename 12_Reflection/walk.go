package walk

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// we make assumptions here about the value passed in
	val := reflect.ValueOf(x)

	// iterate through each field in val
	for i := 0; i < val.NumField(); i++ {
		// access Field i
		field := val.Field(i)

		// check the type of Field i
		if field.Kind() == reflect.String {
			fn(field.String())
		}

		if field.Kind() == reflect.Struct {
			walk(field.Interface(), fn)
		}
	}
}
