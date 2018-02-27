package main

type ProjectConfig struct {
	Name			string  `Project Name`
	Path            string  `Project Path`
	Type            int     `Project Type(simple project, multi project, module)`
}

type PJMaker interface {
	Making(config *ProjectConfig)
}