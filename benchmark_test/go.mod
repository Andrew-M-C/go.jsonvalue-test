module github.com/Andrew-M-C/go.jsonvalue-test/benchmark_test

go 1.16

require (
	github.com/Andrew-M-C/go.jsonvalue v1.3.1
	github.com/Andrew-M-C/go.jsonvalue000 v0.0.0
	github.com/Andrew-M-C/go.jsonvalue103 v0.0.0-00010101000000-000000000000
	github.com/Andrew-M-C/go.jsonvalue105 v0.0.0-00010101000000-000000000000
	github.com/Andrew-M-C/go.jsonvalue111 v0.0.0-00010101000000-000000000000
	github.com/buger/jsonparser v1.1.1
	github.com/bytedance/sonic v1.4.0
	github.com/json-iterator/go v1.1.12
	github.com/mailru/easyjson v0.7.7
)

replace (
	github.com/Andrew-M-C/go.jsonvalue000 => ../../go.jsonvalue
	github.com/Andrew-M-C/go.jsonvalue103 => github.com/Andrew-M-C/go.jsonvalue v1.0.3
	github.com/Andrew-M-C/go.jsonvalue105 => github.com/Andrew-M-C/go.jsonvalue v1.0.5
	github.com/Andrew-M-C/go.jsonvalue111 => github.com/Andrew-M-C/go.jsonvalue v1.1.1
)
