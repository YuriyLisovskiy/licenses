#!/usr/bin/env bash

echo Running golang api tests...
go test -v -timeout 1h -covermode=count -coverprofile=./api/golang/coverage.out ./api/golang
echo Generating coverage report for golang api...
go tool cover -html ./api/golang/coverage.out -o ./api/golang/coverage.html
echo Done.
echo ""
