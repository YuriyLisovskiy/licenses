# Copyright (c) 2018 Yuriy Lisovskiy
# Distributed under the MIT software license, see the accompanying
# file LICENSE or https://opensource.org/licenses/MIT

COVER_OUT = coverage.out

all: clean test

test: go-test py-test

go-test:
	@bash ./scripts/test-go-api.sh

py-test:
	@bash ./scripts/test-py-api.sh

deploy-py-api:
	python3 -m pip install --user --upgrade setuptools wheel
	python3 ./api/py/setup.py sdist bdist_wheel -d ./api/py/dist/
	python3 -m pip install --user --upgrade twine
	twine upload ./api/py/dist/*

clean:
	rm -rf ./build/ ./pyoslapi.egg-info/ ./api/py/dist/ ./dist/
	rm -rf ./api/go/$(COVER_OUT) ./api/go/coverage.html
	rm -rf ./api/py/.coverage ./api/py/coverage/ .coverage