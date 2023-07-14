package benchmark

import (
	"bytes"
	"encoding/json"
	"strconv"
	"testing"

	jsonvalue000 "github.com/Andrew-M-C/go.jsonvalue"
	jsonvalue103 "github.com/Andrew-M-C/go.jsonvalue103"
	jsonvalue105 "github.com/Andrew-M-C/go.jsonvalue105"
	jsonvalue111 "github.com/Andrew-M-C/go.jsonvalue111"
	jsonvalue133 "github.com/Andrew-M-C/go.jsonvalue133"
	jsonparser "github.com/buger/jsonparser"
	sonic "github.com/bytedance/sonic"
	jsoniter "github.com/json-iterator/go"
)

//go:generate go get -u github.com/buger/jsonparser
//go:generate go get -u github.com/bytedance/sonic
//go:generate go get -u github.com/json-iterator/go
//go:generate go get -u github.com/mailru/easyjson
//go:generate go mod tidy
//go:generate easyjson -all object.go
//go:generate go test -bench=. -run=none -benchmem -benchtime=2s
//go:generate go version

// go test -bench=. -run=none -benchmem -benchtime=2s && go version

var unmarshalText = []byte(`{"int":123456,"float":123.456789,"string":"Hello, world!","object":{"int":123456,"float":123.456789,"string":"Hello, world!","object":{"int":123456,"float":123.456789,"string":"Hello, world!","object":{"int":123456,"float":123.456789,"string":"Hello, world!","object":{"int":123456,"float":123.456789,"string":"Hello, world!"},"array":[{"int":123456,"float":123.456789,"string":"Hello, world!"},{"int":123456,"float":123.456789,"string":"Hello, world!"}]}}},"array":[{"int":123456,"float":123.456789,"string":"Hello, world!"},{"int":123456,"float":123.456789,"string":"Hello, world!"}]}`)
var jsonit = jsoniter.ConfigCompatibleWithStandardLibrary

func generateLongObject() []byte {
	buff := bytes.Buffer{}

	buff.WriteRune('{')

	for i := 0; i < 100; i++ {
		if i > 0 {
			buff.WriteRune(',')
		}

		buff.WriteRune('"')
		buff.WriteString(strconv.FormatInt(int64(i), 10))
		buff.WriteString(`":`)
		buff.Write(unmarshalText)
	}

	buff.WriteRune(('}'))

	return buff.Bytes()
}

func init() {
	jsonvalue133.SetDefaultMarshalOptions(jsonvalue133.OptUTF8())
	jsonvalue000.SetDefaultMarshalOptions(jsonvalue000.OptUTF8())
}

func Benchmark_Unmarshal_结构体_json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		o := object{}
		json.Unmarshal(unmarshalText, &o)
	}
}

func Benchmark_Unmarshal_结构体_sonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		o := object{}
		sonic.Unmarshal(unmarshalText, &o)
	}
}

func Benchmark_Unmarshal_结构体_easyjson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		o := object{}
		o.UnmarshalJSON(unmarshalText)
	}
}

func Benchmark_Unmarshal_结构体_manual_by_jsonvalue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		o := object{}
		o.UnmarshalJsonvalue(unmarshalText)
	}
}

func Benchmark_Unmarshal_结构体_jsoniter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		o := object{}
		jsonit.Unmarshal(unmarshalText, &o)
	}
}

func Benchmark_Unmarshal_map_any_json(b *testing.B) {
	raw := unmarshalText
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m := map[string]any{}
		json.Unmarshal(raw, &m)
	}
}

func Benchmark_Unmarshal_map_any_sonic(b *testing.B) {
	raw := unmarshalText
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m := map[string]any{}
		sonic.Unmarshal(raw, &m)
	}
}

func Benchmark_Unmarshal_map_any_jsoniter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := map[string]any{}
		jsonit.Unmarshal(unmarshalText, &m)
	}
}

func Benchmark_Unmarshal_map_any_json_blob(b *testing.B) {
	raw := generateLongObject()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m := map[string]any{}
		json.Unmarshal(raw, &m)
	}
}

func Benchmark_Unmarshal_map_any_jsoniter_blob(b *testing.B) {
	raw := generateLongObject()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m := map[string]any{}
		jsonit.Unmarshal(raw, &m)
	}
}

func Benchmark_Unmarshal_map_any_sonic_blob(b *testing.B) {
	raw := generateLongObject()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m := map[string]any{}
		sonic.Unmarshal(raw, &m)
	}
}

func Benchmark_Unmarshal_any_json(b *testing.B) {
	raw := unmarshalText
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var m any
		json.Unmarshal(raw, &m)
	}
}

func Benchmark_Unmarshal_any_sonic(b *testing.B) {
	raw := unmarshalText
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var m any
		sonic.Unmarshal(raw, &m)
	}
}

func Benchmark_Unmarshal_Jsonvalue_v1_0_3(b *testing.B) {
	origB := unmarshalText
	for i := 0; i < b.N; i++ {
		jsonvalue103.Unmarshal(origB)
	}
}

func Benchmark_Unmarshal_Jsonvalue_v1_0_4(b *testing.B) {
	origB := unmarshalText
	for i := 0; i < b.N; i++ {
		jsonvalue105.Unmarshal(origB)
	}
}

func Benchmark_Unmarshal_Jsonvalue_v1_1_1(b *testing.B) {
	origB := unmarshalText
	for i := 0; i < b.N; i++ {
		jsonvalue111.Unmarshal(origB)
	}
}

func Benchmark_Unmarshal_Jsonvalue_v1_3_3(b *testing.B) {
	origB := unmarshalText
	for i := 0; i < b.N; i++ {
		jsonvalue133.Unmarshal(origB)
	}
}

func Benchmark_Unmarshal_Jsonvalue_develop(b *testing.B) {
	origB := unmarshalText
	for i := 0; i < b.N; i++ {
		jsonvalue000.Unmarshal(origB)
	}
}

func Benchmark__Marshal__map_any_json(b *testing.B) {
	m := map[string]any{}
	json.Unmarshal(unmarshalText, &m)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(&m)
		if err != nil {
			b.Errorf("marshal error: %v", err)
			return
		}
	}
}

func Benchmark__Marshal__结构体_json(b *testing.B) {
	o := object{}
	json.Unmarshal(unmarshalText, &o)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(&o)
		if err != nil {
			b.Errorf("marshal error: %v", err)
			return
		}
	}
}

func Benchmark__Marshal__结构体_sonic(b *testing.B) {
	o := object{}
	sonic.Unmarshal(unmarshalText, &o)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := sonic.Marshal(&o)
		if err != nil {
			b.Errorf("marshal error: %v", err)
			return
		}
	}
}

func Benchmark__Marshal__结构体_jsoniter(b *testing.B) {
	o := object{}
	jsonit.Unmarshal(unmarshalText, &o)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		json.Marshal(&o)
	}
}

func Benchmark__Marshal__结构体_easyjson(b *testing.B) {
	o := object{}
	o.UnmarshalJSON(unmarshalText)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		o.MarshalJSON()
	}
}

func Benchmark__Marshal__结构体_jsonvalue(b *testing.B) {
	o := object{}
	o.UnmarshalJSON(unmarshalText)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		v, _ := jsonvalue133.Import(&o)
		v.MustMarshal()
	}
}

func Benchmark__Import___结构体_jsonvalue(b *testing.B) {
	o := object{}
	o.UnmarshalJSON(unmarshalText)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jsonvalue133.Import(&o)
	}
}

func Benchmark__Import___结构体_jsonvalue_json中转(b *testing.B) {
	o := object{}
	o.UnmarshalJSON(unmarshalText)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b, _ := json.Marshal(&o)
		jsonvalue133.UnmarshalNoCopy(b)
	}
}

func Benchmark__Import___结构体_jsonvalue_sonic中转(b *testing.B) {
	o := object{}
	o.UnmarshalJSON(unmarshalText)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b, _ := sonic.Marshal(&o)
		jsonvalue133.UnmarshalNoCopy(b)
	}
}

func Benchmark__Marshal__Jsoniter_MapItf(b *testing.B) {
	m := map[string]any{}
	jsonit.Unmarshal(unmarshalText, &m)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(&m)
		if err != nil {
			b.Errorf("marshal error: %v", err)
			return
		}
	}
}

func Benchmark____Get____Jsoniter(b *testing.B) {
	raw := unmarshalText
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jsoniter.Get(raw)
	}
}

func Benchmark____Get____Jsoniter_Full(b *testing.B) {
	raw := unmarshalText
	var handleArray func(jsoniter.Any)
	var handleObject func(jsoniter.Any)

	handleArray = func(any jsoniter.Any) {
		le := any.Size()
		for i := 0; i < le; i++ {
			sub := any.Get(i)
			switch sub.ValueType() {
			default:
			// basic type do nothing
			case jsoniter.ArrayValue:
				handleArray(sub)
			case jsoniter.ObjectValue:
				handleObject(sub)
			}
		}
	}

	handleObject = func(any jsoniter.Any) {
		keys := any.Keys()
		for _, k := range keys {
			sub := any.Get(k)
			switch sub.ValueType() {
			default:
			// basic type do nothing
			case jsoniter.ArrayValue:
				handleArray(sub)
			case jsoniter.ObjectValue:
				handleObject(sub)
			}
		}
	}

	for i := 0; i < b.N; i++ {
		any := jsoniter.Get(raw)
		switch any.ValueType() {
		default:
			// basic type do nothing
		case jsoniter.ArrayValue:
			handleArray(any)
		case jsoniter.ObjectValue:
			handleObject(any)
		}
	}
}

func Benchmark____Get____Jsoniter_AndGetParsedValue(b *testing.B) {
	raw := unmarshalText
	any := jsoniter.Get(raw)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		any.Get("object", "object", "object", "array", 1)
	}
}

func Benchmark____Get____Jsoniter_ReadOneChain(b *testing.B) {
	raw := unmarshalText
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		any := jsoniter.Get(raw)
		any.Get("object", "object", "object", "array", 1)
	}
}

func Benchmark____Get____Jsoniter_ReadLevelOne(b *testing.B) {
	raw := unmarshalText
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		any := jsoniter.Get(raw)
		any.Keys()
		// keys := any.Keys()
		// for _, k := range keys {
		// 	any.Get(k)
		// }
	}
}

func Benchmark____Get____Jsoniter_ReadOneChain_Blob(b *testing.B) {
	raw := generateLongObject()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jsoniter.Get(raw, "10", "object", "object", "object", "array", 1)
	}
}

func Benchmark_Unmarshal_Jsonparser_Full(b *testing.B) {
	var objEach func([]byte, []byte, jsonparser.ValueType, int) error
	var arrEach func([]byte, jsonparser.ValueType, int, error)

	objEach = func(_, v []byte, t jsonparser.ValueType, _ int) (noErr error) {
		switch t {
		default:
			// do nothing
		case jsonparser.Array:
			jsonparser.ArrayEach(v, arrEach)
		case jsonparser.Object:
			jsonparser.ObjectEach(v, objEach)
		}
		return
	}

	raw := unmarshalText
	b.ResetTimer()

	arrEach = func(v []byte, t jsonparser.ValueType, _ int, _ error) {
		switch t {
		default:
			// do nothing
		case jsonparser.Array:
			jsonparser.ArrayEach(v, arrEach)
		case jsonparser.Object:
			jsonparser.ObjectEach(v, objEach)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jsonparser.ObjectEach(raw, objEach)
	}
}

func Benchmark_Unmarshal_Jsonparser_ReadLevelOne(b *testing.B) {
	objEach := func(k, v []byte, t jsonparser.ValueType, _ int) (noErr error) {
		return
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jsonparser.ObjectEach(unmarshalText, objEach)
	}
}

func Benchmark_Unmarshal_Jsonparser_ReadLevelOne_Blob(b *testing.B) {
	objEach := func(k, v []byte, t jsonparser.ValueType, _ int) (noErr error) {
		return
	}

	raw := generateLongObject()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jsonparser.ObjectEach(raw, objEach)
	}
}

func Benchmark_Unmarshal_Jsonparser_ReadOneChain(b *testing.B) {
	raw := unmarshalText
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jsonparser.Get(raw, "object", "object", "object", "array", "[1]")
	}
}

func Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob(b *testing.B) {
	raw := generateLongObject()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jsonparser.Get(raw, "10", "object", "object", "object", "array", "[1]")
	}
}

func Benchmark_Unmarshal_Jsonvalue_ReadOneChain(b *testing.B) {
	origB := unmarshalText
	for i := 0; i < b.N; i++ {
		v, _ := jsonvalue133.Unmarshal(origB)
		v.Get("object", "object", "object", "array", 1)
	}
}

func Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob(b *testing.B) {
	origB := generateLongObject()
	for i := 0; i < b.N; i++ {
		v, _ := jsonvalue133.Unmarshal(origB)
		v.Get("50", "object", "object", "object", "array", 1)
	}
}

func Benchmark_Unmarshal_Jsonvalue_GetByInterface(b *testing.B) {
	v, _ := jsonvalue133.Unmarshal(unmarshalText)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		v.Get("object", "object", "object", "array", 1)
	}
}

func Benchmark_Unmarshal_Jsonvalue_GetByInterfaceCaseless(b *testing.B) {
	v, _ := jsonvalue133.Unmarshal(unmarshalText)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		v.Caseless().Get("object", "object", "object", "array", 1)
	}
}

func Benchmark_Unmarshal_Jsonvalue_blob(b *testing.B) {
	// origB := unmarshalText
	origB := generateLongObject()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jsonvalue133.Unmarshal(origB)
	}
}

func Benchmark_Unmarshal_Jsonvalue_develop_blob(b *testing.B) {
	// origB := unmarshalText
	origB := generateLongObject()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jsonvalue000.Unmarshal(origB)
	}
}

func Benchmark_Unmarshal_Jsonvalue_NoCopy(b *testing.B) {
	// origB := unmarshalText
	origB := generateLongObject()

	lst := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		bytes := make([]byte, len(origB))
		copy(bytes, origB)
		lst[i] = bytes
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// err := jsonit.Unmarshal(raw, &s)
		jsonvalue133.UnmarshalNoCopy(lst[i])
	}
}

func Benchmark__Marshal__Jsonvalue(b *testing.B) {
	j, _ := jsonvalue133.Unmarshal(unmarshalText)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := j.Marshal()
		if err != nil {
			b.Errorf("marshal error: %v", err)
			return
		}
	}
}

func Benchmark__Marshal__Jsonvalue_develop(b *testing.B) {
	j, _ := jsonvalue000.Unmarshal(unmarshalText)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := j.Marshal()
		if err != nil {
			b.Errorf("marshal error: %v", err)
			return
		}
	}
}
