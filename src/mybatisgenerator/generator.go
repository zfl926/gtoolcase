package main

import (
	"text/template"
	"os"
)

type JavaClassField struct {
	FieldName           string
	JavaType            string
}

type MyBatisPOJO struct {
	PackageName        string
	ClassName          string
	ImportClasses      []string
	Field              []JavaClassField
	GetSetters         []JavaMethod
}

type JavaMethod struct {
	MethodName        string
	MethodReturn      string
	MethodBody        string
	params            []JavaParam
}


type JavaParam struct {
	ParamName         string
	ParamType         string
}

type JavaInterface struct {
	ReturnType         string
	MethodName         string
	Annotations        []string
	Params             []JavaParam
}

type MyBtisDao  struct {
	PackageName             string
	ImportClasses           []string
	ClassName               string
	InterMethods            []JavaInterface
}


func main() {
	// manager, err := GetDBManager("", "", "", "", 3306)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// desc, err := manager.Desc("hello")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(dec)
	
	var dao MyBtisDao
	dao.ImportClasses = make([]string, 0)
	dao.ImportClasses = append(dao.ImportClasses, "hello111")
	dao.ImportClasses = append(dao.ImportClasses, "hello222")
	dao.PackageName = "test"
	dao.ClassName = "MyTest"
	

	tmpl, err:= template.New("MybatisTemplate").Parse(MybatisPOJOTemplate)
	if err == nil {
		tmpl.Execute(os.Stdout, dao)
	}

}