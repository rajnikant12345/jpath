// Package jpath is created to support getting json path value and provide support for modifying value at a perticular JSON path.
//
// [1] for array index 1, we use 1 based index, not 0 based
// [*] for all array index
// . for object access
// and that's it...
// e.g.
// {
//	"firstName": "John",
// 	"lastName": "doe",
// 	"age": 26,
// 	"address": {
// 		"streetAddress": "naist street",
// 		"city": "Nara",
// 		"postalCode": "630-0192"
// 	},
// 	"phoneNumbers": {
// 		"fika": [
// 			[
// 				[{
// 						"type": "iPhone",
// 						"number": "0123-4567-8888"
// 					},
// 					{
// 						"type": "home",
// 						"number": "0123-4567-8910"
// 					}
// 				]
// 			]
// 		]
// 	}
//
// }


// To access firstname, just say firstName
// To access streetAddress, just say address.streetAddress
// To access number in type home:
//    phoneNumbers.fika.[1].[1].[2].number
// To modify just see example below.




//	Example:

//	package main

//	import "github.com/rajnikant12345/jpath.git"
//	import "encoding/json"
//	import "fmt"
//
//	func main() {
//		j, _ := jpath.CompileNewJsonPath([]string{"phoneNumbers.fika.[1].[1].[*].number", "firstName"})
//
//		in := `{
//		"firstName": "John",
//		"lastName": "doe",
//		"age": 26,
//		"address": {
//			"streetAddress": "naist street",
//			"city": "Nara",
//			"postalCode": "630-0192"
//		},
//		"phoneNumbers": {
//			"fika": [
//				[
//					[{
//							"type": "iPhone",
//							"number": "0123-4567-8888"
//						},
//						{
//							"type": "home",
//							"number": "0123-4567-8910"
//						}
//					]
//				]
//			]
//		}
//
//	}`
//		m := map[string]interface{}{}
//
//		json.Unmarshal( []byte(in), &m )
//
//		out,err := j.MapJsonAtPathValue( m, func(in interface{}) interface{} {
//			fmt.Println(in)
//			return "hello"
//		} )
//
//		if err != nil {
//			fmt.Println(err)
//		} else {
//			fmt.Println(out)
//		}
//
//
//	}
//
//
//
package jpath
