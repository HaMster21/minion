package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func SetupRoutes(router *httprouter.Router) {
	router.GET("/", ShowIndex)
}

func ShowIndex(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	fmt.Fprintf(writer, "This is a response")
}
