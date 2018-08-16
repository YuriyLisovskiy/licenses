# Copyright (c) 2018 Yuriy Lisovskiy
# Distributed under the MIT software license, see the accompanying
# file LICENSE or https://opensource.org/licenses/MIT

from unittest import TestCase

from pyoslapi.util import license_data
from pyoslapi.consts import GNU_ORG, OPEN_SOURCE_ORG


class UtilTest(TestCase):
	license_data_test_data = [
		{
			'input': 'bsd-2-clause',
			'expected_name': 'BSD 2-Clause License',
			'expected_link': OPEN_SOURCE_ORG + 'BSD-2-Clause',
		},
		{
			'input': 'bsd-3-clause',
			'expected_name': 'BSD 3-Clause License',
			'expected_link': OPEN_SOURCE_ORG + 'BSD-3-Clause',
		},
		{
			'input': 'agpl-3.0',
			'expected_name': 'GNU Affero General Public License v3.0',
			'expected_link': GNU_ORG + 'agpl-3.0',
		},
		{
			'input': 'gpl-2.0',
			'expected_name': 'GNU General Public License v2.0',
			'expected_link': GNU_ORG + 'gpl-2.0',
		},
		{
			'input': 'gpl-3.0',
			'expected_name': 'GNU General Public License v3.0',
			'expected_link': GNU_ORG + 'gpl-3.0',
		},
		{
			'input': 'lgpl-2.1',
			'expected_name': 'GNU Lesser General Public License v2.1',
			'expected_link': GNU_ORG + 'lgpl-2.1',
		},
		{
			'input': 'lgpl-3.0',
			'expected_name': 'GNU Lesser General Public License v3.0',
			'expected_link': GNU_ORG + 'lgpl-3.0',
		},
		{
			'input': 'apache-2.0',
			'expected_name': 'Apache License Version 2.0',
			'expected_link': OPEN_SOURCE_ORG + 'Apache-2.0',
		},
		{
			'input': 'epl-2.0',
			'expected_name': 'Eclipse Public License - v2.0',
			'expected_link': OPEN_SOURCE_ORG + 'EPL-2.0',
		},
		{
			'input': 'mit',
			'expected_name': 'MIT License',
			'expected_link': OPEN_SOURCE_ORG + 'MIT',
		},
		{
			'input': 'mpl-2.0',
			'expected_name': 'Mozilla Public License Version 2.0',
			'expected_link': OPEN_SOURCE_ORG + 'MPL-2.0',
		},
		{
			'input': 'unlicense',
			'expected_name': 'Unlicense',
			'expected_link': 'http://unlicense.org',
		},
	]
	
	def test_license_data(self):
		for item in self.license_data_test_data:
			actual_name, actual_link = license_data(item['input'])
			self.assertEqual(actual_name, item['expected_name'])
			self.assertEqual(actual_link, item['expected_link'])
	
	def test_license_data_type_error(self):
		self.assertRaises(TypeError, license_data, 10)
		
	def test_license_data_value_error(self):
		self.assertRaises(ValueError, license_data, 'some-license')
