package basic

import (
	"fmt"
	"testing"
)

func Test_Add(t *testing.T) {
	sum := Add(1, 2)
	if sum == 3 {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}
}
func TestAdd1(t *testing.T) {
	t.Error("the result is error")
}

func Add(a, b int) int {
	fmt.Println("I am in Add func")
	return a + b
}
