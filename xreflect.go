package xreflect

import "reflect"

func IsZero(value any) bool {
	if value == nil {
		return true
	}

	valueOf := reflect.ValueOf(value)

	for {
		if valueOf.Kind() == reflect.Interface ||
			valueOf.Kind() == reflect.Pointer {
			valueOf = valueOf.Elem()
		} else {
			break
		}
	}

	return valueOf.IsZero()
}

func IsNil(value any) bool {
	if value == nil {
		return true
	}

	valueOf := reflect.ValueOf(value)

	switch valueOf.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return valueOf.IsNil()
	default:
		return false
	}
}
