package main
  
import (
    "fmt"
    "testing"
)

func TestMain(t *testing.T) {

    var cost int

    cost = e(2, 5, 1, 3, 5)
    fmt.Printf("Expected cost is 60 and obtained cost of E(2,5,1,3,5) is ")
    fmt.Println(cost)

    if cost != 60 {
        t.Errorf("Fail, expected 61 and got %d", cost)
    }

    cost = e(3, 20, 4, 9, 17)
    fmt.Printf("Expected cost is 2358 and obtained cost of E(3, 20, 4, 9, 17) is ")
    fmt.Println(cost)
    
    if cost != 2358 {
        t.Errorf("Fail, expected 2358 and got %d", cost)
    }    
    
}