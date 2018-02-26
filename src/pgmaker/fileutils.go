package main

import (
	"os"
)

/**
* check if the path exits
 */
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

/**
* create file
 */
func Create(path string) (string, error) {
	if IsExist(path) {
		return path, nil
	}
	if err := os.MkdirAll(path, 0777); err != nil {
		return "", err
	}
	return path, nil
}
