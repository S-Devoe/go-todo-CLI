package command

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/S-Devoe/go-todo-cli/todo"
)

type CmdFlags struct {
	Add string
	Del int
	Edit string
	Toggle int
	List bool
}

func NewCmdFlags() *CmdFlags {
	cf := &CmdFlags{}

	flag.StringVar(&cf.Add, "add", "","Add a new todo, specify todo title")
	flag.StringVar(&cf.Edit,"edit","","Edit a todo by index & specify a new title, id:new_title")
	flag.IntVar(&cf.Del,"del",-1,"Speecify a todo by index to delete")
	flag.IntVar(&cf.Toggle,"toggle",-1,"Specify a todo by index to toggle its completion status")
	flag.BoolVar(&cf.List,"list",false,"List all todos")


	flag.Parse()
	return cf
}


func (cf *CmdFlags) Execute(t *todo.Todos){
	switch{
	case cf.List:
		    t.Print()
	case cf.Add != "":
		    t.AddTodo(cf.Add)
	case cf.Edit != "":
		    parts :=  strings.SplitN(cf.Edit, ":", 2)

            if len(parts) != 2 {
				fmt.Println("Error, invalid format for edit, please use id:new_title")
				os.Exit(1)
			}
			index, err := strconv.Atoi(parts[0])

			if err != nil || index < 0 {
                fmt.Println("Error: invalid index for edit")
                os.Exit(1)
            }
			t.EditTodo(index, parts[1])
	case cf.Toggle != -1:
			t.ToggleCompleted(cf.Toggle)
	case cf.Del!= -1:
		t.DeleteTodo(cf.Del)
	default:
		fmt.Println("Invalid command")
	}
}