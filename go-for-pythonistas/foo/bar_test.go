package foo

import (
    "testing"
)

func Test_Bar(t *testing.T) {
    
    if Bar(1, 2) != 3 {
        t.Error("Something when wrong!")
    } else {
        t.Log("Bar test passed")
    }

}