package utils

import (
  "fmt"
  "testing"
)

func TestSum(t *testing.T) {

    fmt.Println("testing begins...")
    cd100 := New(100)

    fmt.Println("-------- 100 pre-cooldown", cd100)
    cd100.CoolDown(100)
    fmt.Println("-------- 100 after ------", cd100)

    if cd100.CheckAndReset() {
       t.Errorf("Timer went off and it shouldn't have...")
    }

    fmt.Println("-------- 100 pre-cooldown", cd100)
    cd100.CoolDown(1)
    fmt.Println("-------- 100 after ------", cd100)

    if !cd100.CheckAndReset() {
       t.Errorf("Timer didn't go off and it should have...")
    }

    if cd100.CheckAndReset() {
       t.Errorf("Timer went off and it shouldn't have...")
    }


    cd2 := New(2);

    fmt.Println("-------- 2 pre-cooldown", cd2)
    cd2.CoolDown(3)
    fmt.Println("-------- 2 after ------", cd2)

    if !cd2.CheckAndReset() {
       t.Errorf("Timer didn't go off and it should have...")
    }

    fmt.Println("-------- 2 pre-cooldown", cd2)
    cd2.CoolDown(1)
    fmt.Println("-------- 2 after ------", cd2)

    if cd2.CheckAndReset() {
        t.Errorf("Timer went off and it shouldn't have...")
    }
}
