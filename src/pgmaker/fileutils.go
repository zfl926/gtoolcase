package main

import (
	"os"
)

type CreateFileFunc func(path string, fileName string) (string, error)
type CreateFoldFunc func(path string) (string, error)

type Repository struct {
	Name              				string     		 `name`
	Path              				string     		 `path`
	RType              				int        		 `1:file 2:folder`
	ParentReposity                  *Repository      `parent reposity`
	SubRepositories     			[]*Repository    `sub reposity`
	CreateFile                      CreateFileFunc   `create file func`
	CreateFold                      CreateFoldFunc   `create folder func`                  
}
/**
* create folders
*/
func (this *Repository) Create() {
	var createNow bool = false
	var currentPath string = this.Path + GetPathSeparator() + this.Name
	if this.ParentReposity == nil {
		createNow = true
	} else {
		var parentPath string = this.ParentReposity.Path + GetPathSeparator() + this.ParentReposity.Name
		if parentPath == this.Path {
			createNow = true
		}
	}
	// 如果上一级目标和当前路径一直，就开始新建文件或则目录
	if createNow {
		if this.RType == 2 {
			if this.CreateFold != nil {
				this.CreateFold(currentPath)
			}
		} else {
			if  this.CreateFile != nil { 
				this.CreateFile(this.Path, this.Name)
			}
		}
	}
	for _,subRepos := range this.SubRepositories {
		if subRepos != nil {
			subRepos.Create()
		}
	}
}

func GetPathSeparator() string {
	return string(os.PathSeparator)
}
/**
* check if the path exits
*/
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
/**
* create folder
*/
func CreateFolder(path string) (string, error) {
	if IsExist(path) {
		return path, nil
	}
	if err := os.MkdirAll(path, 0777); err != nil {
		return "", err
	}
	return path, nil
}
/**
*  create file
*/
func CreateFile(path string) (string, error){
	return path, nil
}