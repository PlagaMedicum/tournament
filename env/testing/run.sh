#!/bin/sh

echo Installing packages...
go install tournament

sleep 1

echo Testing application...
export CWD=$PWD
go test -v -race ./...

echo 0