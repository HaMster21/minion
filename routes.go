package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type Task struct {
	Description string
}

var allTasks []Task = []Task{Task{"Do work with me"}, Task{"Buy some milk"}, Task {"Get things done"}}

func SetupRoutes(router *httprouter.Router) {
	router.GET("/", ShowIndex)

	router.GET("/tasks", showTasks)
	router.GET("/tasks/:id", showTaskByID)
}

func ShowIndex(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	fmt.Fprintf(writer, "This is a response")
}

func showTasks(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	for _, task := range allTasks {
		fmt.Fprintf(writer, "%s\n", task.Description)
	}
}

func showTaskByID(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, _ := strconv.Atoi(params.ByName("id"))
	fmt.Fprintf(writer, "%s", allTasks[id - 1])
}
