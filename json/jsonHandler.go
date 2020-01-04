package main

import (
	"encoding/json"
	"log"
	"fmt"
	"net/http"
)

type Message struct {
	Name string
	Body string
	Time int64
}

//encode JSON data
// https://golang.org/pkg/encoding/json/
// https://blog.golang.org/json-and-go
func handlerJSON(w http.ResponseWriter, r *http.Request) {
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, _ := json.Marshal(m)
	fmt.Fprintf(w, "%s", b)
}

//decode JSON data
func handlerStringfy(w http.ResponseWriter, r *http.Request) {
	m1 := Message{"Alice", "Hello", 1294706395881547000}
	b, _ := json.Marshal(m1)
	var m Message
	err := json.Unmarshal(b, &m)
	if err != nil {
	}

	fmt.Fprintf(w, "stringfy: 'Name':'%s', 'Body':'%s'", m.Name, m.Body)
}

//decode arbitrary JSON data
func handlerArbitraryStringfy(w http.ResponseWriter, r *http.Request) {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err == nil {
		m := f.(map[string]interface{})

		for k, v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Fprintf(w, "%s (string) : %s\n", k, v)
			case float64:
				fmt.Fprintf(w, "%s (float64): %v\n", k, float64(vv))
			case []interface{}:
				fmt.Fprintf(w, "%s (array): %s\n", k, v)
				for _, u := range vv {
					fmt.Fprintf(w, "%s has %s\n", k, u)
					// fmt.Println(i)
				}
			default:
				fmt.Fprintf(w, "%s is of a type I don't know how to handle\n", k)
			}
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/json", handlerJSON)
	http.HandleFunc("/stringfy", handlerStringfy)
	http.HandleFunc("/arbitrarystringfy", handlerArbitraryStringfy)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
