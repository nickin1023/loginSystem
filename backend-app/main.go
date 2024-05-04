package main

import (
	"backend-app/util"
)

func main() {
	// http.HandleFunc("/", echoHello)
	// http.ListenAndServe(":8000", nil)
	util.FindAll()
}

// func echoHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "<h1>Hello World</h1>")
// }
