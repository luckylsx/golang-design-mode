package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println(color.RedString("$ cd %s", "name"))
}
