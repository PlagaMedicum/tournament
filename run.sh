#!/bin/sh

echo Installing packages[if needed]...
go get -d -v ./...

echo Running application...

go run ./main.go

echo 0