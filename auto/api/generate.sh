#!/bin/sh

protoc -I . -I ./../../src -I /Users/mihail/lib/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.9.1/ --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --validate_out="lang=go:." api.proto
