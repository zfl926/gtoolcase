package main


var MybatisJavaTemplate string = 
`package {{.PackageName}}.dao;

public interface {{.Name}}Dao {
	{{.InterfaceMethod}}
}
`

var MybatisPOJOTemplate string =
`
package {{.PackageName}}.entity;

public interface {{.Name}} {
	{{.PojoProperties}}
}
`