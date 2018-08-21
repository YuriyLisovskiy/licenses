# Copyright (c) 2018 Yuriy Lisovskiy
# Distributed under the MIT software license, see the accompanying
# file LICENSE or https://opensource.org/licenses/MIT

import unittest

from . import unittest as test

test_modules = [
	test.__name__ + '.test_util',
	test.__name__ + '.test_client',
	test.__name__ + '.test_license'
]


def create_test_suite():
	suite = unittest.TestSuite()
	for test_module in test_modules:
		# load all the test cases from the module.
		suite.addTest(unittest.defaultTestLoader.loadTestsFromName(test_module))
	return suite
