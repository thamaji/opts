package opts

import (
	"encoding/json"
	"fmt"
)

func FromUint16Pointer(value *uint16) Uint16 {
	return Uint16{value: value}
}

func NewUint16(value uint16) Uint16 {
	return Uint16{value: &value}
}

func NewEmptyUint16() Uint16 {
	return Uint16{value: nil}
}

type Uint16 struct {
	value *uint16
}

func (self Uint16) IsEmpty() bool {
	return self.value == nil
}

func (self Uint16) Slice() []uint16 {
	if self.value == nil {
		return nil
	}
	return []uint16{*(self.value)}
}

func (self Uint16) Get() (uint16, bool) {
	if self.value == nil {
		return 0, false
	}
	return *(self.value), true
}

func (self Uint16) GetOrDefault(value uint16) uint16 {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Uint16) GetOrElse(f func() uint16) uint16 {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Uint16) ToPointer() *uint16 {
	return self.value
}

func (self Uint16) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Uint16) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Uint16) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
