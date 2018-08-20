# Copyright (c) 2018 Yuriy Lisovskiy
# Distributed under the MIT software license, see the accompanying
# file LICENSE or https://opensource.org/licenses/MIT

COVER_OUT = coverage.out

all: clean test

test: go-test py-test

go-test:
	bash ./scripts/test-go-api.sh

py-test:
	bash ./scripts/test-py-api.sh

deploy-py-api:
	python3 -m pip install --user --upgrade setuptools wheel
	python3 ./api/py-oslapi/setup.py sdist bdist_wheel -d ./api/py-oslapi/dist/
	python3 -m pip install --user --upgrade twine
	twine upload ./api/py-oslapi/dist/*

clean:
	rm -rf ./build/ ./pyoslapi.egg-info/ ./api/py-oslapi/dist/ ./dist/
	rm -rf ./api/go-oslapi/$(COVER_OUT) ./api/go-oslapi/coverage.html
	rm -rf ./api/py-oslapi/.coverage ./api/py-oslapi/coverage/ .coverage