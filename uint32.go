package opts

import (
	"encoding/json"
	"fmt"
)

func FromUint32Pointer(value *uint32) Uint32 {
	return Uint32{value: value}
}

func NewUint32(value uint32) Uint32 {
	return Uint32{value: &value}
}

func NewEmptyUint32() Uint32 {
	return Uint32{value: nil}
}

type Uint32 struct {
	value *uint32
}

func (self Uint32) IsEmpty() bool {
	return self.value == nil
}

func (self Uint32) Slice() []uint32 {
	if self.value == nil {
		return nil
	}
	return []uint32{*(self.value)}
}

func (self Uint32) Get() (uint32, bool) {
	if self.value == nil {
		return 0, false
	}
	return *(self.value), true
}

func (self Uint32) GetOrDefault(value uint32) uint32 {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Uint32) GetOrElse(f func() uint32) uint32 {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Uint32) ToPointer() *uint32 {
	return self.value
}

func (self Uint32) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Uint32) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Uint32) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
