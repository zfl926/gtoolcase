package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input your project name:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There ware errors reading,exiting program.")
		return
	}
	l := len(input)
	projectName := input[0:(l - 1)]
	if projectName != "" {
		fmt.Println("please input path:")
		input, err = inputReader.ReadString('\n')
		l = len(input)
		pathName := input[0:(l - 1)]
		if err != nil {
			fmt.Println("There ware errors reading,exiting program.")
			return
		}
		mavnConfig := CreateWebMavenConfig(projectName, "test", pathName, "pckagename")
		mavnConfig.Making()
	} else {
		return
	}
	
}
