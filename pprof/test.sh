#!/bin/bash

go get -u github.com/json-iterator/go
go get -u github.com/Andrew-M-C/go.jsonvalue
go mod tidy
go run .
echo "go tool pprof -http=:6060 ./jsonvalue-unmarshal.profile"
echo "go tool trace jsonvalue-unmarshal-trace.profile"
go version