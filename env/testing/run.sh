#!/bin/sh

echo Installing packages...

go get ./...
# go get golang.org/x/tools/cmd/cover

echo Testing application...

go test ./pkg/controllers/api/routers/tournament
go test ./pkg/controllers/api/routers/user
# go test ./pkg/controllers/repositories/postgresql/tournament -cwd="/go/src/tournament"
go test ./pkg/controllers/repositories/postgresql/user -cwd="/go/src/tournament"
go test ./pkg/usecases/tournament
go test ./pkg/usecases/user

echo 0