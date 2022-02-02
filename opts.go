package opts

import (
	"encoding/json"
	"fmt"
)

func New[T any](v T) Opt[T] {
	return Opt[T]{v}
}

func Empty[T any]() Opt[T] {
	return Opt[T]{}
}

type Opt[T any] []T

func (opt Opt[T]) IsEmpty() bool {
	return len(opt) == 0
}

func (opt Opt[T]) Get() (T, bool) {
	if len(opt) == 0 {
		return *new(T), false
	}
	return opt[0], true
}

func (opt Opt[T]) GetOrElse(v T) T {
	if len(opt) == 0 {
		return v
	}
	return opt[0]
}

func (opt Opt[T]) Ptr() *T {
	if len(opt) == 0 {
		return nil
	}
	return &(opt[0])
}

func (opt Opt[T]) String() string {
	if len(opt) == 0 {
		return "<nil>"
	}
	return fmt.Sprint(opt[0])
}

func (opt Opt[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(opt.Ptr())
}

func (opt *Opt[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*opt = []T{}
		return nil
	}

	v := *new(T)
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*opt = []T{v}
	return nil
}
