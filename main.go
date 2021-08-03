package main

import (
	"acp10/configs"
	"acp10/routes"
)

func main() {
	configs.InitDB()
	e := routes.New()
	e.Start(":8000")
<<<<<<< HEAD
	// A
=======
	// List
>>>>>>> origin/list
}
