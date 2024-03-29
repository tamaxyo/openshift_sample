package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func main() {
	http.HandleFunc("/", hello)
	bind := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	fmt.Printf("listening on %s...", bind)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello, world from %s on %s:%s", runtime.Version(), os.Getenv("HOST"), os.Getenv("PORT"))
}
