package binding

import "reflect"

type UnmarshalSetter interface {
	UnmarshalSetter(data []byte, setter reflect.Value) error
}
