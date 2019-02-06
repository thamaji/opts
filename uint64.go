package opts

import (
	"encoding/json"
	"fmt"
)

func FromUint64Pointer(value *uint64) Uint64 {
	return Uint64{value: value}
}

func NewUint64(value uint64) Uint64 {
	return Uint64{value: &value}
}

func NewEmptyUint64() Uint64 {
	return Uint64{value: nil}
}

type Uint64 struct {
	value *uint64
}

func (self Uint64) IsEmpty() bool {
	return self.value == nil
}

func (self Uint64) Slice() []uint64 {
	if self.value == nil {
		return nil
	}
	return []uint64{*(self.value)}
}

func (self Uint64) Get() (uint64, bool) {
	if self.value == nil {
		return 0, false
	}
	return *(self.value), true
}

func (self Uint64) GetOrDefault(value uint64) uint64 {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Uint64) GetOrElse(f func() uint64) uint64 {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Uint64) ToPointer() *uint64 {
	return self.value
}

func (self Uint64) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Uint64) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Uint64) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
