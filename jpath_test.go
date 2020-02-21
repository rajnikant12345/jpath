package jpath

import (
	"testing"
	"encoding/json"
	"fmt"
)

func TestGetJsonAtPath(t *testing.T) {
	j, _ := CompileNewJsonPath([]string{"phoneNumbers.fika.[1].[1].[*].number"})

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

	out,err := j.GetJsonAtPathValue( m )

	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}


func TestJsonPath_MapJsonAtPathValue(t *testing.T) {
	j, _ := CompileNewJsonPath([]string{"phoneNumbers.fika.[1].[1].[*].number", "firstName"})

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
		t.Fatal(err)
	}
	t.Log(out)

}
