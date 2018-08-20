### Golang Open Source Licenses API

### Installation
```
$ go get github.com/YuriyLisovskiy/licenses/api/go
```
### Usage

Example:
```go
package main

import (
	"fmt"
	
	"github.com/YuriyLisovskiy/licenses/api/go"
)

func main()  {
	client := oslapi.Client{}
	license, err := client.GetLicense("apache-2.0")
	if err != nil {
		panic(err)
	}
	fmt.Printf(
		"Name: %s\nLink: %s\nContent:\n%s\n",
		license.Title(), license.Link(), license.Content(),
	)
	
	header, err := client.GetHeader("apache-2.0")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Header: %s", header)
	
	list, err := client.GetList()
	if err != nil {
		panic(err)
	}
	fmt.Printf("List: %s", list)
}
```
