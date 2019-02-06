package opts

import (
	"encoding/json"
	"fmt"
)

func FromInt8Pointer(value *int8) Int8 {
	return Int8{value: value}
}

func NewInt8(value int8) Int8 {
	return Int8{value: &value}
}

func NewEmptyInt8() Int8 {
	return Int8{value: nil}
}

type Int8 struct {
	value *int8
}

func (self Int8) IsEmpty() bool {
	return self.value == nil
}

func (self Int8) Slice() []int8 {
	if self.value == nil {
		return nil
	}
	return []int8{*(self.value)}
}

func (self Int8) Get() (int8, bool) {
	if self.value == nil {
		return 0, false
	}
	return *(self.value), true
}

func (self Int8) GetOrDefault(value int8) int8 {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Int8) GetOrElse(f func() int8) int8 {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Int8) ToPointer() *int8 {
	return self.value
}

func (self Int8) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Int8) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Int8) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
