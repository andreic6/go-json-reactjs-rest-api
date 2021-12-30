package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)
type myData struct {
	Test string `json:"send"`
}
	func handle(w http.ResponseWriter,  r *http.Request) {
         w.Header().Set("Access-Control-Allow-Origin", "*")
	 w.Header().Set("Access-Control-Allow-Methods","POST")
	 var data myData
	 decoder := json.NewDecoder(r.Body)
	 err := decoder.Decode(&data)
	 if err != nil {
	     fmt.Println("eroare")
	 }
	 defer r.Body.Close()
	 value := data.Test
	 fmt.Fprintf(w,"%s", value)
}
func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe("192.168.1.1:3000", nil)
}


