# Copyright (c) 2018 Yuriy Lisovskiy
# Distributed under the MIT software license, see the accompanying
# file LICENSE or https://opensource.org/licenses/MIT

import unittest

from tests.test_runner import create_test_suite


if __name__ == '__main__':
	unittest.TextTestRunner().run(create_test_suite())
