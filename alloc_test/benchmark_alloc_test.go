package alloc

import (
	"encoding/json"
	"testing"

	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
)

// go test -bench=. -run=none -benchmem -benchtime=0.1s

func Benchmark_初始化一个_float(b *testing.B) {
	raw := []byte("123.456")
	for i := 0; i < b.N; i++ {
		jsonvalue.MustUnmarshalNoCopy(raw)
	}
}

func Benchmark_初始化一个_int(b *testing.B) {
	raw := []byte("-123")
	for i := 0; i < b.N; i++ {
		jsonvalue.MustUnmarshalNoCopy(raw)
	}
}

func Benchmark_初始化一个_uint(b *testing.B) {
	raw := []byte("123")
	for i := 0; i < b.N; i++ {
		jsonvalue.MustUnmarshalNoCopy(raw)
	}
}

func Benchmark_初始化一个_bool(b *testing.B) {
	raw := []byte("true")
	for i := 0; i < b.N; i++ {
		jsonvalue.MustUnmarshalNoCopy(raw)
	}
}

func Benchmark_初始化一个_array_空的(b *testing.B) {
	raw := []byte(`[]`)
	for i := 0; i < b.N; i++ {
		jsonvalue.MustUnmarshalNoCopy(raw)
	}
}

func Benchmark_初始化一个_array_1个元素(b *testing.B) {
	raw := []byte(`[1]`)
	for i := 0; i < b.N; i++ {
		jsonvalue.MustUnmarshalNoCopy(raw)
	}
}

func Benchmark_初始化一个_array_2个元素(b *testing.B) {
	raw := []byte(`[1,1]`)
	for i := 0; i < b.N; i++ {
		jsonvalue.MustUnmarshalNoCopy(raw)
	}
}

func Benchmark_初始化一个_array_32个元素(b *testing.B) {
	raw := []byte(`[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]`)
	for i := 0; i < b.N; i++ {
		jsonvalue.MustUnmarshalNoCopy(raw)
	}
}

func Benchmark_初始化一个_array_33个元素(b *testing.B) {
	raw := []byte(`[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]`)
	for i := 0; i < b.N; i++ {
		jsonvalue.MustUnmarshalNoCopy(raw)
	}
}

func Benchmark_初始化一个_array_34个元素(b *testing.B) {
	raw := []byte(`[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]`)
	for i := 0; i < b.N; i++ {
		jsonvalue.MustUnmarshalNoCopy(raw)
	}
}

func Benchmark_初始化一个_array_35个元素(b *testing.B) {
	raw := []byte(`[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]`)
	for i := 0; i < b.N; i++ {
		jsonvalue.MustUnmarshalNoCopy(raw)
	}
}

func Benchmark_初始化一个_object_空的(b *testing.B) {
	raw := []byte(`{}`)
	for i := 0; i < b.N; i++ {
		jsonvalue.MustUnmarshalNoCopy(raw)
	}
}

func Benchmark_初始化一个_object_1个元素(b *testing.B) {
	raw := []byte(`{"1":1}`)
	for i := 0; i < b.N; i++ {
		jsonvalue.MustUnmarshalNoCopy(raw)
	}
}

func Benchmark_初始化一个_object_2个元素(b *testing.B) {
	raw := []byte(`{"1":1,"2":2}`)
	for i := 0; i < b.N; i++ {
		jsonvalue.MustUnmarshalNoCopy(raw)
	}
}

func Benchmark_初始化一个_object_2个元素_原生_json_map(b *testing.B) {
	raw := []byte(`{"1":1,"2":2}`)
	for i := 0; i < b.N; i++ {
		var m map[string]interface{}
		json.Unmarshal(raw, &m)
	}
}

func Benchmark_初始化一个_object_2个元素_原生_json_struct(b *testing.B) {
	raw := []byte(`{"1":1,"2":2}`)
	for i := 0; i < b.N; i++ {
		var m struct {
			One int `json:"1"`
			Two int `json:"2"`
		}
		json.Unmarshal(raw, &m)
	}
}

func Benchmark_初始化一个_string(b *testing.B) {
	raw := []byte(`"string"`)
	for i := 0; i < b.N; i++ {
		jsonvalue.MustUnmarshalNoCopy(raw)
	}
}
