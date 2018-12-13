package main


var MybatisJavaTemplate string = 
`
package {{.PackageName}}.dao;

{{ range .ImportClasses }}
	import {{.}};
{{end}}


public interface {{.ClassName}}Dao {
	{{ range .InterMethods }}
		{{ range .Annotations }}
		{{.}}
		{{end}}
		{{.ReturnType}} {{.MethodName}} ({{ range .Params }}{{.ParamName}} {{.ParamType}}, {{end}});
	{{end}}
}
`

var MybatisPOJOTemplate string =
`
package {{.PackageName}}.entity;

{{ range .ImportClasses }}
	import {{.}};
{{end}}


public class {{.ClassName}} {
	{{ range .Field }}
	private {{.JavaType}} {{.FieldName}}
	{{end}}

	{{ range .GetSetters }}
	public {{.MethodReturn}} {{.MethodName}} ({{ range .Params }}{{.ParamName}} {{.ParamType}}, {{end}}){
		{{.MethodBody}}
	}
	{{end}}
}
`