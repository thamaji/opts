package opts

import (
	"encoding/json"
	"fmt"
)

func FromRunePointer(value *rune) Rune {
	return Rune{value: value}
}

func NewRune(value rune) Rune {
	return Rune{value: &value}
}

func NewEmptyRune() Rune {
	return Rune{value: nil}
}

type Rune struct {
	value *rune
}

func (self Rune) IsEmpty() bool {
	return self.value == nil
}

func (self Rune) Slice() []rune {
	if self.value == nil {
		return nil
	}
	return []rune{*(self.value)}
}

func (self Rune) Get() (rune, bool) {
	if self.value == nil {
		return 0, false
	}
	return *(self.value), true
}

func (self Rune) GetOrDefault(value rune) rune {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self Rune) GetOrElse(f func() rune) rune {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self Rune) ToPointer() *rune {
	return self.value
}

func (self Rune) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self Rune) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *Rune) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
