package opts

import (
	"encoding/json"
	"fmt"
)

func FromInt16Pointer(value *int16) Int16 {
	return Int16{value: value}
}

func NewInt16(value int16) Int16 {
	return Int16{value: &value}
}

func NewEmptyInt16() Int16 {
	return Int16{value: nil}
}

type Int16 struct {
	value *int16
}

func (self Int16) IsEmpty() bool {
	return self.value == nil
}

func (self Int16) Slice() []int16 {
	if self.value == nil {
		return nil
	}
	return []int16{*(self.value)}
}

func (self Int16) Get() (int16, bool) {
	if self.value == nil {
		return 0, false
	}
	return *(self.value), true
}

func (self Int16) GetOrDefault(value int16) int16 {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Int16) GetOrElse(f func() int16) int16 {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Int16) ToPointer() *int16 {
	return self.value
}

func (self Int16) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Int16) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Int16) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
