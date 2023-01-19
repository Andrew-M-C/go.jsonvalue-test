#!/bin/bash

go get -u github.com/buger/jsonparser
go get -u github.com/bytedance/sonic
go get -u github.com/json-iterator/go
go get -u github.com/mailru/easyjson
go get -u github.com/Andrew-M-C/go.jsonvalue
go mod tidy
easyjson -all object.go
go test -bench=. -run=none -benchmem -benchtime=2s
go version