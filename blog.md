## The JSON Path Library

While working on Data Protection Gateway ( DPG ), we got stuck on a problem. The problem was, to tokenize JSON tags. This problem has two subparts.

* To read which token to modify.
* Modify the token.

DPG as a product needs to function as a **Bump On The Wire**, or reverse proxy i.e. 
* Read the incoming HTTP request.
* Tokenize based on content type.
* Send a request to upstream server and return a response.

We can write a separate blog on DPG but, this is about how we solved the JSON modification problem.

 Let's start by taking a JSON request as an example, and we want to modify all  type inside phoneNumbers.
{
	"firstName": "John",
	"lastName": "doe",
	"age": 26,
	"address": {
		"streetAddress": "naist street",
		"city": "Nara",
		"postalCode": "630-0192"
	},
	"phoneNumbers": {
		"fika": 
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
	}
}

So, the two problems are as follows:
How to access type inside the phoneNumbers object.
Modify each value of the selected tag i.e. type

Solution (https://github.com/rajnikant12345/jpath)

The open-source library given in the link is developed by me to solve the mentioned problem. I will be explaining in detail how it works and solves the problem.

Part1 ( Access JSON objects) :

To access JSON objects, one needs to know JSON Path Syntax. You can refer to https://jsonpath.com/ for testing it. In my example JSON request, phoneNumbers.fika.[*].[*].type is the JSON path to access all type values in phoneNumbers object.  
Parsing the JSON path is a big deal of problem, I tried my best but, timelines were not in favor of writing a JSON path parser from scratch, so I used Golang Viper ( https://github.com/spf13/viper ), to parse JSON path.
 It is not fully compatible with JSON Path Syntax but, it does the minimum job to parse it. Viper is not case sensitive, so I took its code, modified it and, used it as JSON Path parser. Viper gave me a map of maps as output, which helped me to look deep inside JSON objects. So, now I can access a JSON tag at any level and modify its value as needed. So, for phoneNumbers.fika.[*].[*].type, the map entry will look like,

map[phoneNumber]:{map[fika]:{ map[*]:{map[*]:{map[type]:{}} }   }}

each key in map key corresponds to a level of JSON object. So, phoneNumbers at
 level 0, * at level 1, another star at level 2 and finally the object type . Now we have put this JSON path into a structure and now it's time to do the job modifying it inside a JSON data.

Part2 ( Modify JSON Data ) :
The second part of the problem was easy to solve but, to implement one need to know, Depth First Search and Golang reflection. I parsed the data using standard Golang JSON parser, and it converted my JSON data to a map of interfaces, e.g.

map[phoneNumber]:{map[fika]:[ [ map[type]:"iphone",map[number]:"0123-4567-8888" ],[  map[home]:"iphone",map[number]:"0123-4567-1234" ]   ]  }}
  
So, the logic is simple, explore the depth of your JSON Path and check your JSON at the same level. To modify the JSON value, you have to register a callback. 
The best part is Golang JSON Parser decode JSON request and save it in to map of interfaces, so we don't need to write our JSON parser.

Given below is an example code for using the library. 
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





Conclusion
The code is working fine, but there are still some things to do:
Remove Viper and add code to parse complete JSON Path Syntax.
We can be more efficient while parsing JSON.

Thanks for reading and happy coding.
