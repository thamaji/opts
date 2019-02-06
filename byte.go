package opts

import (
	"encoding/json"
	"fmt"
)

func FromBytePointer(value *byte) Byte {
	return Byte{value: value}
}

func NewByte(value byte) Byte {
	return Byte{value: &value}
}

func NewEmptyByte() Byte {
	return Byte{value: nil}
}

type Byte struct {
	value *byte
}

func (self Byte) IsEmpty() bool {
	return self.value == nil
}

func (self Byte) Slice() []byte {
	if self.value == nil {
		return nil
	}
	return []byte{*(self.value)}
}

func (self Byte) Get() (byte, bool) {
	if self.value == nil {
		return 0, false
	}
	return *(self.value), true
}

func (self Byte) GetOrDefault(value byte) byte {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Byte) GetOrElse(f func() byte) byte {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Byte) ToPointer() *byte {
	return self.value
}

func (self Byte) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Byte) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Byte) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
