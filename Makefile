# Copyright (c) 2018 Yuriy Lisovskiy
# Distributed under the MIT software license, see the accompanying
# file LICENSE or https://opensource.org/licenses/MIT

COVER_OUT = coverage.out

all: clean test

test: test-golang test-python

test-golang:
	@echo Running golang api tests...
	@go test -v -timeout 1h -covermode=count -coverprofile=./api/golang/$(COVER_OUT) ./api/golang
	@echo Generating coverage report for golang api...
	@go tool cover -html ./api/golang/$(COVER_OUT) -o ./api/golang/coverage.html
	@echo Done

test-python:
	@echo Running python api tests...
	@coverage run ./api/python/tests.py
	@coverage report
	@echo Generating coverage report for python api...
	@coverage html -d ./api/python/coverage/
	@echo Done.

clean:
	@rm -rf ./api/golang/$(COVER_OUT) ./api/golang/coverage.html
	@rm -rf ./api/python/.coverage ./api/python/coverage/ .coverage