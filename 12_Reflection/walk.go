package walk

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// we make assumptions here about the value passed in
	val := reflect.ValueOf(x)

	// iterate through each field in val
	for i := 0; i < val.NumField(); i++ {
		// access Field i
		field := val.Field(i)
		// call the String method. May not actually have one if it's not a string
		fn(field.String())
	}
}
