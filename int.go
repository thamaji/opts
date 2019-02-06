package opts

import (
	"encoding/json"
	"fmt"
)

func FromIntPointer(value *int) Int {
	return Int{value: value}
}

func NewInt(value int) Int {
	return Int{value: &value}
}

func NewEmptyInt() Int {
	return Int{value: nil}
}

type Int struct {
	value *int
}

func (self Int) IsEmpty() bool {
	return self.value == nil
}

func (self Int) Slice() []int {
	if self.value == nil {
		return nil
	}
	return []int{*(self.value)}
}

func (self Int) Get() (int, bool) {
	if self.value == nil {
		return 0, false
	}
	return *(self.value), true
}

func (self Int) GetOrDefault(value int) int {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Int) GetOrElse(f func() int) int {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Int) ToPointer() *int {
	return self.value
}

func (self Int) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Int) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
