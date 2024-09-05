package main
import "fmt"

import (
	"go-todo-api/database"
	"go-todo-api/env"
	"go-todo-api/router"
)

func main() {
	env.Init()
	database.Init()
	router.Init()
        fmt.print("Adding a print statement")
}
