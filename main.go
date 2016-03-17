package main

import (
	"fmt"
	"github.com/hamster21/minion/task"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	stdin := fmt.Sprintf("%s", data)
	fmt.Printf("Data from stdin: %s\n", stdin)

	task, err := task.FromJSON(stdin)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Resulting task object:\n%v\n", task)
	fmt.Println("Hello world, this is a work in progress :)")
}
