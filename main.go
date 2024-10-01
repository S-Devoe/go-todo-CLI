package main

import (
	"github.com/S-Devoe/go-todo-cli/command"
	"github.com/S-Devoe/go-todo-cli/storage"
	"github.com/S-Devoe/go-todo-cli/todo"
)

func main() {
	t := todo.Todos{}
	s := storage.NewStorage[todo.Todos]("todos.json")

	s.LoadFile(&t)

	cmdFlag :=  command.NewCmdFlags()
	cmdFlag.Execute(&t)

	s.SaveFile(t)

	


}