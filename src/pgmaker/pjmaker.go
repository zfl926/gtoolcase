package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input your name:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There ware errors reading,exiting program.")
		return
	}
	l := len(input)
	p := input[0:(l - 1)]
	var exits bool = IsExist(p)
	fmt.Print(exits)
	CreateFile(p)
}
