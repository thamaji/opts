package opts

import (
	"encoding/json"
	"fmt"
)

func FromFloat32Pointer(value *float32) Float32 {
	return Float32{value: value}
}

func NewFloat32(value float32) Float32 {
	return Float32{value: &value}
}

func NewEmptyFloat32() Float32 {
	return Float32{value: nil}
}

type Float32 struct {
	value *float32
}

func (self Float32) IsEmpty() bool {
	return self.value == nil
}

func (self Float32) Slice() []float32 {
	if self.value == nil {
		return nil
	}
	return []float32{*(self.value)}
}

func (self Float32) Get() (float32, bool) {
	if self.value == nil {
		return 0, false
	}
	return *(self.value), true
}

func (self Float32) GetOrDefault(value float32) float32 {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Float32) GetOrElse(f func() float32) float32 {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Float32) ToPointer() *float32 {
	return self.value
}

func (self Float32) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Float32) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Float32) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
