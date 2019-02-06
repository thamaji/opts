package opts

import (
	"encoding/json"
	"fmt"
)

func FromInt32Pointer(value *int32) Int32 {
	return Int32{value: value}
}

func NewInt32(value int32) Int32 {
	return Int32{value: &value}
}

func NewEmptyInt32() Int32 {
	return Int32{value: nil}
}

type Int32 struct {
	value *int32
}

func (self Int32) IsEmpty() bool {
	return self.value == nil
}

func (self Int32) Slice() []int32 {
	if self.value == nil {
		return nil
	}
	return []int32{*(self.value)}
}

func (self Int32) Get() (int32, bool) {
	if self.value == nil {
		return 0, false
	}
	return *(self.value), true
}

func (self Int32) GetOrDefault(value int32) int32 {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Int32) GetOrElse(f func() int32) int32 {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Int32) ToPointer() *int32 {
	return self.value
}

func (self Int32) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Int32) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Int32) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
