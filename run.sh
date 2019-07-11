#!/bin/sh

echo Installing packages[if needed]...
go get ./...

echo Running application...

go run ./main.go

echo 0