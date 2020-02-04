package main

import (
	"encoding/json"
	"fmt"
	"go-example/stringutil"
)

func fib(x int) int64 {
	var f1, f2 int64 = 0,1

	if x<=1 {
		return int64(x)
	} else {
		for i :=0; i<x; i++ {
			f1, f2 = f2, f1+f2
		}
		
		return f1
	}
}

//decode arbitrary JSON data
func handlerArbitraryStringfy() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err == nil {
		m := f.(map[string]interface{})

		for k, v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case float64:
				fmt.Println(k, "is float64", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}
	}
}

func main() {
	fmt.Printf("hello World!!\n")
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	handlerArbitraryStringfy()
	fmt.Println(fib(6))
}
