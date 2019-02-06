package opts

import (
	"encoding/json"
	"fmt"
)

func FromInterfacePointer(value *interface{}) Interface {
	return Interface{value: value}
}

func NewInterface(value interface{}) Interface {
	return Interface{value: &value}
}

func NewEmptyInterface() Interface {
	return Interface{value: nil}
}

type Interface struct {
	value *interface{}
}

func (self Interface) IsEmpty() bool {
	return self.value == nil
}

func (self Interface) Slice() []interface{} {
	if self.value == nil {
		return nil
	}
	return []interface{}{*(self.value)}
}

func (self Interface) Get() (interface{}, bool) {
	if self.value == nil {
		return nil, false
	}
	return *(self.value), true
}

func (self Interface) GetOrDefault(value interface{}) interface{} {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Interface) GetOrElse(f func() interface{}) interface{} {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Interface) ToPointer() *interface{} {
	return self.value
}

func (self Interface) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Interface) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Interface) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
