### Golang Open Source Licenses API

### Installation
```
$ go get github.com/YuriyLisovskiy/licenses/api/golang
```
### Usage

An example how to get a license from https://github.com/YuriyLisovskiy/licenses:

```go
package main

import (
	"fmt"
	
	"github.com/YuriyLisovskiy/licenses/api/golang"
)

func main()  {
	client := golang.Client{}
	license, err := client.GetLicense("apache-2.0")
	if err != nil {
		panic(err)
	}
	fmt.Printf(
		"Name: %s\nLink: %s\nContent:\n%s\n",
		license.Name(), license.Link(), license.Content(),
	)
}
```
