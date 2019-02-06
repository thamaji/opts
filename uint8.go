package opts

import (
	"encoding/json"
	"fmt"
)

func FromUint8Pointer(value *uint8) Uint8 {
	return Uint8{value: value}
}

func NewUint8(value uint8) Uint8 {
	return Uint8{value: &value}
}

func NewEmptyUint8() Uint8 {
	return Uint8{value: nil}
}

type Uint8 struct {
	value *uint8
}

func (self Uint8) IsEmpty() bool {
	return self.value == nil
}

func (self Uint8) Slice() []uint8 {
	if self.value == nil {
		return nil
	}
	return []uint8{*(self.value)}
}

func (self Uint8) Get() (uint8, bool) {
	if self.value == nil {
		return 0, false
	}
	return *(self.value), true
}

func (self Uint8) GetOrDefault(value uint8) uint8 {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Uint8) GetOrElse(f func() uint8) uint8 {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Uint8) ToPointer() *uint8 {
	return self.value
}

func (self Uint8) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Uint8) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Uint8) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
