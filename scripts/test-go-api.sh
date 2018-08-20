#!/usr/bin/env bash

echo Running golang api tests...
go test -v -timeout 1h -covermode=count -coverprofile=./api/go/coverage.out ./api/go
echo Generating coverage report for golang api...
go tool cover -html ./api/go/coverage.out -o ./api/go/coverage.html
echo Done.
echo ""
