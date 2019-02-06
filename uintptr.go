package opts

import (
	"encoding/json"
	"fmt"
)

func FromUintptrPointer(value *uintptr) Uintptr {
	return Uintptr{value: value}
}

func NewUintptr(value uintptr) Uintptr {
	return Uintptr{value: &value}
}

func NewEmptyUintptr() Uintptr {
	return Uintptr{value: nil}
}

type Uintptr struct {
	value *uintptr
}

func (self Uintptr) IsEmpty() bool {
	return self.value == nil
}

func (self Uintptr) Slice() []uintptr {
	if self.value == nil {
		return nil
	}
	return []uintptr{*(self.value)}
}

func (self Uintptr) Get() (uintptr, bool) {
	if self.value == nil {
		return 0, false
	}
	return *(self.value), true
}

func (self Uintptr) GetOrDefault(value uintptr) uintptr {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Uintptr) GetOrElse(f func() uintptr) uintptr {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Uintptr) ToPointer() *uintptr {
	return self.value
}

func (self Uintptr) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Uintptr) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Uintptr) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
