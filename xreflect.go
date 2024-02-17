package xreflect

import "reflect"

func IsZero(v any) bool {
	e := reflect.ValueOf(&v)

	for {
		if e.Kind() == reflect.Interface || e.Kind() == reflect.Pointer {
			e = e.Elem()
		} else {
			break
		}
	}

	zero := e.IsZero()

	return zero
}

func IsNil(v any) bool {
	if v == nil {
		return true
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(v).IsNil()
	default:
		return false
	}
}
