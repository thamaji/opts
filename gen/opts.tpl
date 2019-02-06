package opts

import (
	"encoding/json"
	"fmt"
)

func From{{.Type}}Pointer(value *{{.RawType}}) {{.Type}} {
	return {{.Type}}{value: value}
}

func New{{.Type}}(value {{.RawType}}) {{.Type}} {
	return {{.Type}}{value: &value}
}

func NewEmpty{{.Type}}() {{.Type}} {
	return {{.Type}}{value: nil}
}

type {{.Type}} struct {
	value *{{.RawType}}
}

func (self {{.Type}}) IsEmpty() bool {
	return self.value == nil
}

func (self {{.Type}}) Slice() []{{.RawType}} {
	if self.value == nil {
		return nil
	}
	return []{{.RawType}}{*(self.value)}
}

func (self {{.Type}}) Get() ({{.RawType}}, bool) {
	if self.value == nil {
		return {{.ZeroValue}}, false
	}
	return *(self.value), true
}

func (self {{.Type}}) GetOrDefault(value {{.RawType}}) {{.RawType}} {
	if self.value == nil {
		return value
	}
	return *(self.value)
}

func (self {{.Type}}) GetOrElse(f func() {{.RawType}}) {{.RawType}} {
	if self.value == nil {
		return f()
	}
	return *(self.value)
}

func (self {{.Type}}) ToPointer() *{{.RawType}} {
	return self.value
}

func (self {{.Type}}) String() string {
	if self.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*(self.value))
}

func (self {{.Type}}) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.value)
}

func (self *{{.Type}}) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(self.value))
}
