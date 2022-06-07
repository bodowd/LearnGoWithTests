package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numberOfValues := 0

	// this helps us extract something out of the type
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.Struct:
		numberOfValues = val.NumField()
		// if it's a Struct we need to extract out the Field
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		// if it's a Slice, we need to extract out the Index
		getField = val.Index
	case reflect.String:
		fn(val.String())
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	}

	// if it's a Struct or a Slice, we iterate over its values by calling walk on each one

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	// can't use NumField on a pointer Value, so we need to extract the underlying value before we can
	// use NumField in the for loop
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
