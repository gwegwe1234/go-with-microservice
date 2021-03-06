package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

func main()  {
	port := 8081
	http.HandleFunc("/jsonEncoding", helloWorldHandler)

	log.Printf("Server starting on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request!", http.StatusBadRequest)
		return
	}

	var request helloWorldRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Bad Request!", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "hello" + request.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
