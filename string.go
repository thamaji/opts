package opts

import (
	"encoding/json"
	"fmt"
)

func FromStringPointer(value *string) String {
	return String{value: value}
}

func NewString(value string) String {
	return String{value: &value}
}

func NewEmptyString() String {
	return String{value: nil}
}

type String struct {
	value *string
}

func (self String) IsEmpty() bool {
	return self.value == nil
}

func (self String) Slice() []string {
	if self.value == nil {
		return nil
	}
	return []string{*(self.value)}
}

func (self String) Get() (string, bool) {
	if self.value == nil {
		return "", false
	}
	return *(self.value), true
}

func (self String) GetOrDefault(value string) string {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self String) GetOrElse(f func() string) string {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self String) ToPointer() *string {
	return self.value
}

func (self String) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self String) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *String) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
