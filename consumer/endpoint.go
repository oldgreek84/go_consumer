package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w)
}

type Event struct {
	eventType string `json:"_eventType"`
}

var Events Event

func headers(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("server: %s /\n", r.Method)
	fmt.Printf("server: query id: %s\n", r.URL.Query().Get("id"))
	fmt.Printf("server: content-type: %s\n", r.Header.Get("content-type"))
	fmt.Printf("server: headers:\n")
	for headerName, headerValue := range r.Header {
		fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("server: could not read request body: %s\n", err)
	}

  var s string
  var data map[string]interface{}
  json.Unmarshal(reqBody, &s)
  json.Unmarshal([]byte(s), &data)


	fmt.Printf("server: request body: %T\n", reqBody)
	fmt.Printf("server: request body: %s\n", reqBody)
  // fmt.Printf("server: data: %s\n", data["_eventType"])
  // fmt.Printf("server: data: %s\n", data)
  // fmt.Printf("server: data: %T\n", data)
  for k, v := range data {
    fmt.Printf("\n%s => %v", k, v)
  }
  return

}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	// mux := http.NewServeMux()
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8000", nil)
}
