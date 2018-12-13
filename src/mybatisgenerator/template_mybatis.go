package main


var MybatisJavaTemplate string = 
`
package {{.PackageName}}.dao;

{{ range .ImportClasses }}
	import {{.}}
{{end}}


public interface {{.ClassName}}Dao {
	{{ range .InterMethods }}
		{{ range .Annotations }}
		{{.}}
		{{end}}
		{{.ReturnType}} {{.MethodName}} ({{ range .Params }}{{.ParamName}} {{.ParamType}}, {{end}})
	{{end}}
}
`

var MybatisPOJOTemplate string =
`
package {{.PackageName}}.entity;

public interface {{.Name}} {
	{{.PojoProperties}}
}
`