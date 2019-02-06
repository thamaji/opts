package opts

import (
	"encoding/json"
	"fmt"
)

func FromFloat64Pointer(value *float64) Float64 {
	return Float64{value: value}
}

func NewFloat64(value float64) Float64 {
	return Float64{value: &value}
}

func NewEmptyFloat64() Float64 {
	return Float64{value: nil}
}

type Float64 struct {
	value *float64
}

func (self Float64) IsEmpty() bool {
	return self.value == nil
}

func (self Float64) Slice() []float64 {
	if self.value == nil {
		return nil
	}
	return []float64{*(self.value)}
}

func (self Float64) Get() (float64, bool) {
	if self.value == nil {
		return 0, false
	}
	return *(self.value), true
}

func (self Float64) GetOrDefault(value float64) float64 {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Float64) GetOrElse(f func() float64) float64 {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Float64) ToPointer() *float64 {
	return self.value
}

func (self Float64) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Float64) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Float64) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
