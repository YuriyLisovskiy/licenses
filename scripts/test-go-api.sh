#!/usr/bin/env bash

echo Running golang api tests...
go test -v -timeout 1h -covermode=count -coverprofile=./api/go-oslapi/coverage.out ./api/go-oslapi
echo Generating coverage report for golang api...
go tool cover -html ./api/go-oslapi/coverage.out -o ./api/go-oslapi/coverage.html
echo Done.
echo ""
