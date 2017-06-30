// package reflectx contains a set of reflection utilities and well-known types.
package reflectx

import (
	"context"
	"fmt"
	"reflect"
)

// Well-known reflected types. Convenience definitions.
var (
	Bool   = reflect.TypeOf((*bool)(nil)).Elem()
	Int    = reflect.TypeOf((*int)(nil)).Elem()
	Int8   = reflect.TypeOf((*int8)(nil)).Elem()
	Int16  = reflect.TypeOf((*int16)(nil)).Elem()
	Int32  = reflect.TypeOf((*int32)(nil)).Elem()
	Int64  = reflect.TypeOf((*int64)(nil)).Elem()
	Uint   = reflect.TypeOf((*uint)(nil)).Elem()
	Uint8  = reflect.TypeOf((*uint8)(nil)).Elem()
	Uint16 = reflect.TypeOf((*uint16)(nil)).Elem()
	Uint32 = reflect.TypeOf((*uint32)(nil)).Elem()
	Uint64 = reflect.TypeOf((*uint64)(nil)).Elem()
	String = reflect.TypeOf((*string)(nil)).Elem()
	Error  = reflect.TypeOf((*error)(nil)).Elem()

	Context   = reflect.TypeOf((*context.Context)(nil)).Elem()
	Type      = reflect.TypeOf((*reflect.Type)(nil)).Elem()
	ByteSlice = reflect.TypeOf((*[]byte)(nil)).Elem()
)

// IsNumber returns true iff the given type is an integer, float, or complex.
func IsNumber(t reflect.Type) bool {
	return IsInteger(t) || IsFloat(t) || IsComplex(t)
}

// IsInteger returns true iff the given type is an integer, such as
// uint16, int, or int64.
func IsInteger(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	default:
		return false
	}
}

// IsFloat returns true iff the given type is float32 or float64.
func IsFloat(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

// IsComplex returns true iff the given type is complex64 or complex128.
func IsComplex(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Complex64, reflect.Complex128:
		return true
	default:
		return false
	}
}

// SkipPtr returns the target of a Ptr type, if a Ptr. Otherwise itself.
func SkipPtr(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}

// MakeSlice creates a slice of type []T with the given elements.
func MakeSlice(t reflect.Type, values ...reflect.Value) reflect.Value {
	ret := reflect.MakeSlice(reflect.SliceOf(t), len(values), len(values))
	for i, value := range values {
		if value.Type() != t {
			panic(fmt.Sprintf("element type is %v, want %v", value.Type(), t))
		}
		ret.Index(i).Set(value)
	}
	return ret
}
