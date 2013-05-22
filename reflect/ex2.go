package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x int = 4
    v := reflect.ValueOf(&x)
    fmt.Println("type:", v.Type())
    fmt.Println("settability:", v.CanSet())
    fmt.Println("value:", v.Interface())
    fmt.Println("indirect value:", reflect.Indirect(v).Interface())
    p := v.Elem()
    fmt.Println("settability:", p.CanSet())
    p.SetInt(10)
    fmt.Println("new value:", p.Interface())
}
