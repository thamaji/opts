package opts

import (
	"encoding/json"
	"fmt"
)

func FromComplex64Pointer(value *complex64) Complex64 {
	return Complex64{value: value}
}

func NewComplex64(value complex64) Complex64 {
	return Complex64{value: &value}
}

func NewEmptyComplex64() Complex64 {
	return Complex64{value: nil}
}

type Complex64 struct {
	value *complex64
}

func (self Complex64) IsEmpty() bool {
	return self.value == nil
}

func (self Complex64) Slice() []complex64 {
	if self.value == nil {
		return nil
	}
	return []complex64{*(self.value)}
}

func (self Complex64) Get() (complex64, bool) {
	if self.value == nil {
		return 0, false
	}
	return *(self.value), true
}

func (self Complex64) GetOrDefault(value complex64) complex64 {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Complex64) GetOrElse(f func() complex64) complex64 {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Complex64) ToPointer() *complex64 {
	return self.value
}

func (self Complex64) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Complex64) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Complex64) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
