### JavaScript Open Source Licenses API

An API which allows you to get license templates or license notices from your
JavaScript code.

### Installation
```bash
$ npm install oslapi
```

### Usage

Example:
```js
let oslapi = require('oslapi');
    
let client = new oslapi.Client();
    
client.setListHandler(function (list) {
	console.log(list);
}).getList();
    
client.setHeaderHandler(function (header) {
	console.log(header);
}).getHeader('gpl-3.0');
    
client.setLicenseHandler(function (license) {
	console.log(license.title());
	console.log(license.link());
	console.log(license.content());
}).getLicense('mit');
```
