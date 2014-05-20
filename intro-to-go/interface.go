package main

import "fmt"

type Walker interface {
	Walk()
}

type Person struct{}

func (p *Person) Walk() {
	fmt.Println("I'm a person walking")
}

func MakeWalk(w Walker) {
	w.Walk()
}

func main() {
	person := Person{}
	MakeWalk(person)
}
