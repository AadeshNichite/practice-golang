package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
	req.Body.Close()
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
			fmt.Println("Hello")
		}

	}
	req.Body.Close()
}

func main() {

	err := http.ListenAndServe(":8090", nil)
	fmt.Println(err)

	http.HandleFunc("/", hello)
	http.HandleFunc("/headers", headers)

	// if c {
	// 	fmt.Println("need to check")
	// } else {
	// 	fmt.Println("sucess", c)
	// }

}
