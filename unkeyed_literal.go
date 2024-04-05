package main

import "fmt"

type Config struct {
	status bool
	port   int
	host   string
	url    string
	// _      struct{} // un-comment this to prevent unkeyed literals
}

func main() {
	a := Config{
		true, 123, "host", "",
	}

	b := Config{
		status: true,
	}

	fmt.Println(a)
	fmt.Println(b)

}
