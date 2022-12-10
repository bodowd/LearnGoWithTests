package walk

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// we make assumptions here about the value passed in
	val := getValue(x)

	// iterate through each field in val
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	// can't use NumField on a pointer Value
	// need to extract the underlying value before we can do that by using Elem()
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
