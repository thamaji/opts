package opts

import (
	"encoding/json"
	"fmt"
)

func FromInt64Pointer(value *int64) Int64 {
	return Int64{value: value}
}

func NewInt64(value int64) Int64 {
	return Int64{value: &value}
}

func NewEmptyInt64() Int64 {
	return Int64{value: nil}
}

type Int64 struct {
	value *int64
}

func (self Int64) IsEmpty() bool {
	return self.value == nil
}

func (self Int64) Slice() []int64 {
	if self.value == nil {
		return nil
	}
	return []int64{*(self.value)}
}

func (self Int64) Get() (int64, bool) {
	if self.value == nil {
		return 0, false
	}
	return *(self.value), true
}

func (self Int64) GetOrDefault(value int64) int64 {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Int64) GetOrElse(f func() int64) int64 {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Int64) ToPointer() *int64 {
	return self.value
}

func (self Int64) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Int64) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
