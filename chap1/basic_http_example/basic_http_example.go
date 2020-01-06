package basic_http_example

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	port := 8081
	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", 8081)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello World!\n")
}
