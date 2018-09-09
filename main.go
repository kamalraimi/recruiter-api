package main

import "fmt"

func Hello() string {
	return "Hello world"
}

func main() {
	output := Hello()
	fmt.Println(output)
}
