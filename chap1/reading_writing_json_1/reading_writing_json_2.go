package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	// 출력 필드를 message 로 바꾼다
	Message string `json:"message"`

	// 이 필드를 출력하지 않는다. - 사용
	Author string `json:"-"`

	// 값이 비어 있으면 출력하지 않는다.
	Date string `json:",omitempty"`

	// 출력을 문자열로 변환하고 이름을 "id"로 바꾼다.
	Id int `json:"id,string"`
}

func main()  {
	port := 8081
	http.HandleFunc("/jsonEncoding", helloWorldHandlerWithJson)

	log.Printf("Server starting on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port), nil))
}

func helloWorldHandlerWithJson(w http.ResponseWriter, r *http.Request)  {
	response := helloWorldResponse{Message: "Hello World!", Author: "ted.park", Id : 3}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}
