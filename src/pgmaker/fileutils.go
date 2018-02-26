package main

import (
	"os"
	"fmt"
)

/**
* check if the path exits
*/
func IsExist(path string) bool {
	_, err := os.Stat(path)
	fmt.Println(err)
	return err == nil || os.IsExist(err)
}
/**
*  create file
*/
// func Create(dir string, name string) (error){

// }