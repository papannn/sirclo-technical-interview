package main

import (
	"fmt"
	"net/http"

	"github.com/backend/server"
)

func main() {
	mux, err := server.Init()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Running Server at :3000")
	http.ListenAndServe(":3000", mux)
}
