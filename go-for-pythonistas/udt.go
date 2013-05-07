package main

import "fmt"

//An object in go is called a struct
type Person struct {
    Firstname string
    Lastname string
    Phone string
}

//you can attach methods to a struct
func (p *Person) Info() string {
    return fmt.Sprintf("%v %v (%v)", p.Firstname, p.Lastname, p.Phone)
}

func main() {
    p := Person{"Dan", "Worth", "800-555-5555"}
    fmt.Println("Hello", p.Info())
}