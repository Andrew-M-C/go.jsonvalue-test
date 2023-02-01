package benchmark

//go:generate go get -u github.com/mailru/easyjson
//go:generate easyjson -all object.go

type object struct {
	Int    int       `json:"int"`
	Float  float64   `json:"float"`
	String string    `json:"string"`
	Object *object   `json:"object,omitempty"`
	Array  []*object `json:"array,omitempty"`
}
