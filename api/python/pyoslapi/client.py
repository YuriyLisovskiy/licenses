# Copyright (c) 2018 Yuriy Lisovskiy
# Distributed under the MIT software license, see the accompanying
# file LICENSE or https://opensource.org/licenses/MIT

import json
import base64
import requests

from .consts import BASE_URL
from .license import License
from .util import license_data
from .errors import TYPE_ERROR, LICENSE_NOT_FOUND


# Represents simple client for getting a license.
class Client:
	
	# Downloads license from https://github.com/YuriyLisovskiy/licenses repository.
	@staticmethod
	def get_license(name):
		
		# Check if name is of type string.
		if not isinstance(name, str):
			raise TYPE_ERROR
		
		# Download license from https://github.com/YuriyLisovskiy/licenses/licenses.
		response = requests.get(BASE_URL + '/' + name)
		if response.status_code == 200:
			
			# Decode response.
			json_data = json.loads(response.text)
			content = json_data.get('content')
			if content is not None:
				
				# Decode license content using base64 encoding.
				decoded_data = base64.b64decode(content)
				if decoded_data is not None:
					
					# Get license name and source link by its keyword.
					l_name, l_link = license_data(name)
					return License(name=l_name, link=l_link, content=decoded_data.decode('utf-8'))
			return None
		elif response.status_code == 404:
			raise LICENSE_NOT_FOUND
