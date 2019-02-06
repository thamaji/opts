package opts

import (
	"encoding/json"
	"fmt"
)

func FromBoolPointer(value *bool) Bool {
	return Bool{value: value}
}

func NewBool(value bool) Bool {
	return Bool{value: &value}
}

func NewEmptyBool() Bool {
	return Bool{value: nil}
}

type Bool struct {
	value *bool
}

func (self Bool) IsEmpty() bool {
	return self.value == nil
}

func (self Bool) Slice() []bool {
	if self.value == nil {
		return nil
	}
	return []bool{*(self.value)}
}

func (self Bool) Get() (bool, bool) {
	if self.value == nil {
		return false, false
	}
	return *(self.value), true
}

func (self Bool) GetOrDefault(value bool) bool {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Bool) GetOrElse(f func() bool) bool {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Bool) ToPointer() *bool {
	return self.value
}

func (self Bool) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Bool) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
