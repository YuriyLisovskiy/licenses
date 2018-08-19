# Copyright (c) 2018 Yuriy Lisovskiy
# Distributed under the MIT software license, see the accompanying
# file LICENSE or https://opensource.org/licenses/MIT

COVER_OUT = coverage.out

all: clean test

test: gotest pytest

gotest:
	@bash ./scripts/test-go.sh

pytest:
	@bash ./scripts/test-py.sh

clean:
	@rm -rf ./api/golang/$(COVER_OUT) ./api/golang/coverage.html
	@rm -rf ./api/python/.coverage ./api/python/coverage/ .coverage