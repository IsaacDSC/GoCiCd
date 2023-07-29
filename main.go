package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("ALIVE SERVER"))
	})
	fmt.Println("Stated Server on port 3000...")
	if err := http.ListenAndServe(":3000", nil); err == nil {
		panic(err)
	}
}
