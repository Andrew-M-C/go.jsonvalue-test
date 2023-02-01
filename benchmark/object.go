package benchmark

import (
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
)

//go:generate go get -u github.com/mailru/easyjson
//go:generate easyjson -all object.go

type object struct {
	Int    int       `json:"int"`
	Float  float64   `json:"float"`
	String string    `json:"string"`
	Object *object   `json:"object,omitempty"`
	Array  []*object `json:"array,omitempty"`
}

func (o *object) UnmarshalJsonvalue(b []byte) error {
	v, err := jsonvalue.Unmarshal(b)
	if err != nil {
		return err
	}

	o.unmarshalFromJsonvalue(v)
	return nil
}

func (o *object) unmarshalFromJsonvalue(v *jsonvalue.V) {
	o.Int, _ = v.GetInt("int")
	o.Float, _ = v.GetFloat64("float")
	o.String, _ = v.GetString("string")

	if sub, err := v.GetObject("object"); err == nil {
		o.Object = &object{}
		o.Object.unmarshalFromJsonvalue(sub)
	}

	if sub, err := v.GetArray("array"); err == nil {
		o.Array = make([]*object, 0, sub.Len())
		sub.RangeArray(func(_ int, subSub *jsonvalue.V) bool {
			subO := &object{}
			subO.unmarshalFromJsonvalue(subSub)
			o.Array = append(o.Array, subO)
			return true
		})
	}
}
