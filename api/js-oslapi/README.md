### JavaScript Open Source Licenses API

An API which allows you to get license templates or license notices from your
JavaScript code.

### Installation
// TODO: add installation guide.

### Usage

Get list of available licenses:
```javascript
let client = new Client();
client.setListHandler(function (list) {
    console.log(list);
}).getList();
```

Get header:
```javascript
let client = new Client();
client.setHeaderHandler(function (header) {
    console.log(header);
}).getHeader('gpl-3.0');
```

Get license:
```javascript
let client = new Client();
client.setLicenseHandler(function (license) {
    console.log(license.name());
    console.log(license.link());
    console.log(license.content());
}).getLicense('mit');
```