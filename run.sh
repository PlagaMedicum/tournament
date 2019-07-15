#!/bin/sh

echo Installing packages[if needed]...
go get ./...

echo Testing application...

go test ./tests/

echo Running application...

go run ./main.go

echo 0