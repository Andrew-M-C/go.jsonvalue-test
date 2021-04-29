module go.jsonvalue-test

go 1.16

require (
	github.com/Andrew-M-C/go.jsonvalue v1.0.6-0.20210428104912-9fd5ac3bf83a
	github.com/Andrew-M-C/go.jsonvalue103 v0.0.0-00010101000000-000000000000
	github.com/Andrew-M-C/go.jsonvalue105 v0.0.0-00010101000000-000000000000
	github.com/buger/jsonparser v1.1.1
	github.com/json-iterator/go v1.1.10
	github.com/mailru/easyjson v0.7.7
)

replace (
	github.com/Andrew-M-C/go.jsonvalue103 => github.com/Andrew-M-C/go.jsonvalue v1.0.3
	github.com/Andrew-M-C/go.jsonvalue105 => github.com/Andrew-M-C/go.jsonvalue v1.0.5
)
