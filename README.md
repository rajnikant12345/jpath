[![Go Report Card](https://goreportcard.com/badge/github.com/rajnikant12345/jpath)](https://goreportcard.com/report/github.com/rajnikant12345/jpath)
# jpath 
## A small utility library to play with JSON path.

* It is using [viper](https://github.com/spf13/viper) to parse jsonpath.
* It has very easy syntax.
* Can access any value from any path.
* Can set value at any json path.
* Uses go modules.

## How to import
```
 GO111MODULE=on go get github.com/rajnikant12345/jpath
 GO111MODULE=on go mod tidy
```
## How to use

```Go
package main

import "github.com/rajnikant12345/jpath.git"
import "encoding/json"
import "fmt"

func main() {
	j, _ := jpath.CompileNewJsonPath([]string{"phoneNumbers.fika.[1].[1].[*].number", "firstName"})

	in := `{
	"firstName": "John",
	"lastName": "doe",
	"age": 26,
	"address": {
		"streetAddress": "naist street",
		"city": "Nara",
		"postalCode": "630-0192"
	},
	"phoneNumbers": {
		"fika": [
			[
				[{
						"type": "iPhone",
						"number": "0123-4567-8888"
					},
					{
						"type": "home",
						"number": "0123-4567-8910"
					}
				]
			]
		]
	}

}`
	m := map[string]interface{}{}

	json.Unmarshal( []byte(in), &m )

	out,err := j.MapJsonAtPathValue( m, func(in interface{}) interface{} {
		fmt.Println(in)
		return "hello"
	} )

	if err != nil {
		fmt.Println(err)
	} else {
	    fmt.Println(out)
	}


}
```
