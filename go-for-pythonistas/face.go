package main

import (
    "fmt"
    "errors"
)

type Friender interface {
    Friend (string) (bool, error)
}

type Facebook struct {}

func (f Facebook) Friend(name string) (bool, error) {
    if name == "Dan" {
        return true, nil
    }
    return false, errors.New("Facebook only allows friends with the name 'Dan'")
}

func IsFriend(fdr Friender, name string) bool {   
    if friend, err := fdr.Friend(name); err != nil {
        fmt.Println(err.Error()); return friend
    }
    return true
}

func main() {
    f := Facebook{}
    fmt.Println("Is Dan a Friend", IsFriend(f, "Dan"))
    fmt.Println("Is Stan a Friend", IsFriend(f, "Stan"))
}
