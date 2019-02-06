package opts

import (
	"encoding/json"
	"fmt"
)

func FromComplex128Pointer(value *complex128) Complex128 {
	return Complex128{value: value}
}

func NewComplex128(value complex128) Complex128 {
	return Complex128{value: &value}
}

func NewEmptyComplex128() Complex128 {
	return Complex128{value: nil}
}

type Complex128 struct {
	value *complex128
}

func (self Complex128) IsEmpty() bool {
	return self.value == nil
}

func (self Complex128) Slice() []complex128 {
	if self.value == nil {
		return nil
	}
	return []complex128{*(self.value)}
}

func (self Complex128) Get() (complex128, bool) {
	if self.value == nil {
		return 0, false
	}
	return *(self.value), true
}

func (self Complex128) GetOrDefault(value complex128) complex128 {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Complex128) GetOrElse(f func() complex128) complex128 {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Complex128) ToPointer() *complex128 {
	return self.value
}

func (self Complex128) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Complex128) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Complex128) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
