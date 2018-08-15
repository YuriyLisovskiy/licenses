### Python Open Source Licenses API

### Installation
```
$ pip install pyoslapi
```
> Currently not added to PyPi

### Usage

An example how to get a license from https://github.com/YuriyLisovskiy/licenses:

```py
from pyoslapi.client import Client

if __name__ == '__main__':
	client = Client()
	license = client.get_license('apache-2.0')
	if license is not None:
		print(
			"Name: {}\nLink: {}\nContent:\n{}\n".format(
				license.name, license.link, license.content
			)
		)

```
