package main

type MavenConfig struct {
	PjConfig        ProjectConfig
	Group           string   `group id`
	PKName          string   `package name`
}

// func (this *MavenConfig) Create(config *ProjectConfig) {
// 	if config != nil {
// 		projectName   := config.Name
// 		groupName     := this.Group
// 		pakcageName   := this.PKName
// 		path          := config.Path
// 	}
// }

/* create maven web config */
func CreateWebMavenConfig(artifactId string, groupId string, path string, packageName string)(*MavenConfig){
	return &MavenConfig{
		PjConfig : ProjectConfig{
			Name : artifactId,
			Path : path,
			Type : 1,
		},
		Group  : groupId,
		PKName : packageName,
	}
}