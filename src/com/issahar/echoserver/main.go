package main

import (
	"net/http"
	"log"
	"fmt"
	"strings"
	"encoding/json"
)



func Echo(w http.ResponseWriter, r *http.Request) {
	var reqData = make(map[string]string)

	// handle all the headers
	for key, values := range r.Header {
		reqData[key] = strings.Join(values, ",")
	}

	// handle parameters
	r.ParseForm()
	for key, values := range r.Form {
		reqData[key] = strings.Join(values, ",")
	}

	body, _ := json.Marshal(reqData)
	fmt.Fprintf(w, string(body))
}


func main() {
	http.HandleFunc("/", Echo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
