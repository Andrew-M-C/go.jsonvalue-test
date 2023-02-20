module github.com/Andrew-M-C/go.jsonvalue-test/benchmark

go 1.18

require (
	github.com/Andrew-M-C/go.jsonvalue v1.3.4-0.20230220091712-b2a6d94ad0a3
	github.com/Andrew-M-C/go.jsonvalue103 v0.0.0-00010101000000-000000000000
	github.com/Andrew-M-C/go.jsonvalue105 v0.0.0-00010101000000-000000000000
	github.com/Andrew-M-C/go.jsonvalue111 v0.0.0-00010101000000-000000000000
	github.com/Andrew-M-C/go.jsonvalue133 v0.0.0-00010101000000-000000000000
	github.com/buger/jsonparser v1.1.1
	github.com/bytedance/sonic v1.8.1
	github.com/json-iterator/go v1.1.12
	github.com/mailru/easyjson v0.7.7
)

require (
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	golang.org/x/arch v0.2.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
)

replace (
	github.com/Andrew-M-C/go.jsonvalue000 => ../../go.jsonvalue
	github.com/Andrew-M-C/go.jsonvalue103 => github.com/Andrew-M-C/go.jsonvalue v1.0.3
	github.com/Andrew-M-C/go.jsonvalue105 => github.com/Andrew-M-C/go.jsonvalue v1.0.5
	github.com/Andrew-M-C/go.jsonvalue111 => github.com/Andrew-M-C/go.jsonvalue v1.1.1
	github.com/Andrew-M-C/go.jsonvalue133 => github.com/Andrew-M-C/go.jsonvalue v1.3.3
)
