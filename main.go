package main

import "fmt"

type a struct {
	Name string
}

type b a

func main() {
	c := &b{}
	c.Name = "231"
	fmt.Println(*c)
	println("hello world")
}
