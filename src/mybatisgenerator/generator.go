package main

import (
	"fmt"
)





func main() {
	manager, err := GetDBManager("", "", "", "", 3306)
	if err != nil {
		fmt.Println(err)
	}
	desc, err := manager.Desc("hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dec)
}