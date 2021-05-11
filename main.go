package main

import (
	"./controller"
	"net/http"
)

func main()  {
	mux := controller.Register()
	http.ListenAndServe(":3000", mux)
}
