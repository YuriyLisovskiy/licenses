# Copyright (c) 2018 Yuriy Lisovskiy
# Distributed under the MIT software license, see the accompanying
# file LICENSE or https://opensource.org/licenses/MIT

COVER_OUT = coverage.out

all: clean test

test: test-golang

test-golang:
	@echo Running golang api tests...
	@go test -v -timeout 1h -covermode=count -coverprofile=./api/golang/$(COVER_OUT) ./api/golang
	@echo Generating coverage report for golang api...
	@go tool cover -html ./api/golang/$(COVER_OUT) -o ./api/golang/coverage.html
	@echo Done

clean:
	@rm -rf ./api/golang/$(COVER_OUT) ./api/golang/coverage.html