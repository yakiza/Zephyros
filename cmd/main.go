package main

import (
	"github.com/yakiza/Zephyros/api"
	"net/http"
)

func main() {
	r :=  api.NewRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}

}
