package main

import (
	"fmt"

	"github.com/xushuhui/goal/cmd/goal"
)

func main() {
	err := goal.Execute()
	if err != nil {
		fmt.Println("execute error: ", err.Error())
	}
}
