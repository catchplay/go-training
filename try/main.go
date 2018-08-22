package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// start with a string representation of our JSON data
var input = `
{
    "created_at": "Thu May 31 00:00:01 +0000 2012"
}
`

func main() {
	// our target will be of type map[string]interface{}, which is a
	// pretty generic type that will give us a hashtable whose keys
	// are strings, and whose values are of type interface{}
	var val map[string]interface{}

	if err := json.Unmarshal([]byte(input), val); err != nil {
		panic(err)
	}

	fmt.Println(val)
	for k, v := range val {
		fmt.Println(k, reflect.TypeOf(v))
	}
}