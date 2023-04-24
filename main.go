package main

import (
	"fmt"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "hello world change!")
}

func main() {
	http.HandleFunc("/hello", helloWorldPage)
	err := http.ListenAndServe(":8080", nil)
	fmt.Print("Error abc: ", err)
	//http.ListenAndServe("", nil)
}
