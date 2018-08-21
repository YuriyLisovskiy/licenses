# Copyright (c) 2018 Yuriy Lisovskiy
# Distributed under the MIT software license, see the accompanying
# file LICENSE or https://opensource.org/licenses/MIT

from unittest import TestCase

from pyoslapi.license import License


class TestLicense(TestCase):
	
	def test_license_constructor_errors(self):
		self.assertRaises(TypeError, License, 1, 's', 's')
		self.assertRaises(TypeError, License, 's', 1, 's')
		self.assertRaises(TypeError, License, 's', 's', 1)

	def test_license_properties(self):
		l = License('some-name', 'some-link', 'some-content')
		self.assertEqual(l.name, 'some-name')
		self.assertEqual(l.link, 'some-link')
		self.assertEqual(l.content, 'some-content')
