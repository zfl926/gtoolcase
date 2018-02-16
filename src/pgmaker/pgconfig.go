package main

type ProjectConfig struct {
	Name			string  `Project Name`
	Path            string  `Project Path`
	Type            int     `Project Type(simple project, multi project, module)`
}

type MavenConfig struct {
	ProjectConfig
	Group         string   `group id`
	PKName        string   `package name`
}

/* create maven web config */
func CreateWebMavenConfig(artifactId string, groupId string, path string, packageName string)(*MavenConfig){
	return &MavenConfig{
		ProjectConfig : ProjectConfig{
			Name : artifactId,
			Path : path,
			Type : 1,
		},
		Group : groupId,
		PKName : packageName,
	}
}