package main

import (
	"fmt"
	"reflect"
)

type Example struct {
	field1 string
	field2 string
}

func main() {
	myStruct := Example{field1: "test1", field2: "test1"}
	custom_field := "field1"
	val := reflect.ValueOf(myStruct).FieldByName(custom_field)
	fmt.Println(val)
}
