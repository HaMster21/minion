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

	fmt.Printf("Resulting task object:\n%s\n\n", task)
	fmt.Printf("All data:\n%#v\n", task)
}
