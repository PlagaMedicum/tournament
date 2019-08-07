#!/bin/sh

go install tournament
go build -ldflags "-linkmode external -extldflags -static" -o ./bin/main ./main.go
