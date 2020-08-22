package test

import (
	"fmt"
	"testing"

	
)

func TestA(t *testing.T) {
	
}
func swit(types int) {
	var j int
	switch types {
	case 1:
		j = 1
		fmt.Println("1 start")

	case 2:
		j = 2
		fmt.Println("2 start")

	}
	fmt.Println(j)
	fmt.Println("1 end")
}
