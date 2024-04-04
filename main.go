package main

import (
	"fmt"

	"github.com/tech-pratheesh/learning-with-golang/database"
)

func main() {
	op := database.Operations()
	users, err := op.GetUsers()
	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}
