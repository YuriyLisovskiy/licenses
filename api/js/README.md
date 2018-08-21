### JavaScript Open Source Licenses API

An API which allows you to get license templates or license notices from your
JavaScript code.

### Installation
```bash
$ npm install oslapi
```

### Usage

Example:
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
    <p id="list">A list</p>
    <p id="header">A header</p>
    <p id="license">A license</p>
</body>
<script src="path/to/oslapi.js"></script>
<script>

    let client = new Client();
    
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

</script>
</html>
```
