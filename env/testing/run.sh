#!/bin/sh

echo Installing packages...
go install tournament

echo Testing application...
go test -v -race ./pkg/controllers/api/routers/tournament
go test -v -race ./pkg/controllers/api/routers/user
go test -v -race ./pkg/controllers/repositories/postgresql/tournament -cwd="/go/src/tournament"
go test -v -race ./pkg/controllers/repositories/postgresql/user -cwd="/go/src/tournament"
go test -v -race ./pkg/usecases/tournament
go test -v -race ./pkg/usecases/user

echo 0