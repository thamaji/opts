package opts

import (
	"encoding/json"
	"fmt"
)

func FromUintPointer(value *uint) Uint {
	return Uint{value: value}
}

func NewUint(value uint) Uint {
	return Uint{value: &value}
}

func NewEmptyUint() Uint {
	return Uint{value: nil}
}

type Uint struct {
	value *uint
}

func (self Uint) IsEmpty() bool {
	return self.value == nil
}

func (self Uint) Slice() []uint {
	if self.value == nil {
		return nil
	}
	return []uint{*(self.value)}
}

func (self Uint) Get() (uint, bool) {
	if self.value == nil {
		return 0, false
	}
	return *(self.value), true
}

func (self Uint) GetOrDefault(value uint) uint {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Uint) GetOrElse(f func() uint) uint {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Uint) ToPointer() *uint {
	return self.value
}

func (self Uint) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Uint) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Uint) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
