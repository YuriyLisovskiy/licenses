# Copyright (c) 2018 Yuriy Lisovskiy
# Distributed under the MIT software license, see the accompanying
# file LICENSE or https://opensource.org/licenses/MIT

from .consts import GNU_ORG, OPEN_SOURCE_ORG
from .errors import TYPE_ERROR, LICENSE_NOT_FOUND


# Returns license name and link by given license key.
def license_data(key):
	if not isinstance(key, str):
		raise TYPE_ERROR
	if key == 'bsd-2-clause':
		name = 'BSD 2-Clause License'
		link = OPEN_SOURCE_ORG + 'BSD-2-Clause'
	elif key == 'bsd-3-clause':
		name = 'BSD 3-Clause License'
		link = OPEN_SOURCE_ORG + 'BSD-3-Clause'
	elif key == 'agpl-3.0':
		name = 'GNU Affero General Public License v3.0'
		link = GNU_ORG + 'agpl-3.0'
	elif key == 'gpl-2.0':
		name = 'GNU General Public License v2.0'
		link = GNU_ORG + 'gpl-2.0'
	elif key == 'gpl-3.0':
		name = 'GNU General Public License v3.0'
		link = GNU_ORG + 'gpl-3.0'
	elif key == 'lgpl-2.1':
		name = 'GNU Lesser General Public License v2.1'
		link = GNU_ORG + 'lgpl-2.1'
	elif key == 'lgpl-3.0':
		name = 'GNU Lesser General Public License v3.0'
		link = GNU_ORG + 'lgpl-3.0'
	elif key == 'apache-2.0':
		name = 'Apache License Version 2.0'
		link = OPEN_SOURCE_ORG + 'Apache-2.0'
	elif key == 'epl-2.0':
		name = 'Eclipse Public License - v2.0'
		link = OPEN_SOURCE_ORG + 'EPL-2.0'
	elif key == 'mit':
		name = 'MIT License'
		link = OPEN_SOURCE_ORG + 'MIT'
	elif key == 'mpl-2.0':
		name = 'Mozilla Public License Version 2.0'
		link = OPEN_SOURCE_ORG + 'MPL-2.0'
	elif key == 'unlicense':
		name = 'Unlicense'
		link = 'http://unlicense.org'
	else:
		raise LICENSE_NOT_FOUND
	return name, link
