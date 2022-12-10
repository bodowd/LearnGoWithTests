package walk

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// we make assumptions here about the value passed in
	val := reflect.ValueOf(x)
	// look at the first field. Could be that there are no fields at all, which would cause a panic
	field := val.Field(0)
	// call String() but it might not have this method
	fn(field.String())
}
