package main

import (
	"fmt"
	"bufio"
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
    var exits bool = IsExist(input)
    fmt.Printf("Your input is %s" ,exits)
}