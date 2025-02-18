package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/VechniyRabotnik/GolangToDoList/todo"
)

const (
	todoFile = ".todos.json"
)

func main() {

	add := flag.Bool("add", false, "Добавить новую задачу")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("Test")
		err := todos.Storo(todoFile)

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stdout, "Неизвестная команда")
		os.Exit(0)
	}

}
