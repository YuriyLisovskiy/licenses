#!/usr/bin/env bash

echo Running python api tests...
coverage run ./api/py/tests.py
coverage report
echo Generating coverage report for python api...
coverage html -d ./coverage/
echo Done.
echo ""
