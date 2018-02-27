package main

type MavenConfig struct {
	PjConfig        ProjectConfig
	Group           string   `group id`
	PKName          string   `package name`
}

/*******************************************************************
*   folder make
*/
func (this *MavenConfig) Making() {
	projectName   := this.PjConfig.Name
	//groupName     := this.Group
	//pakcageName   := this.PKName
	path          := this.PjConfig.Path
	// set root folder
	var pathRoot string = path + GetPathSeparator() + projectName
	if !IsExist(pathRoot) {
		var rootReposity *Repository = &Repository{
			Name  : projectName,
			Path  : path,
			RType : 2,
		}
		rootReposity.ParentReposity = nil
		rootReposity.SubRepositories = make([]*Repository, 3)
		var relyReposity *Repository = &Repository {
			Name  : "rely",
			Path  : pathRoot,
			RType : 2,
		}
		relyReposity.ParentReposity = rootReposity
		rootReposity.SubRepositories = append(rootReposity.SubRepositories, relyReposity)
		var serviceReposity *Repository = &Repository {
			Name  : "service",
			Path  : pathRoot,
			RType : 2,
		}
		serviceReposity.ParentReposity = rootReposity
		rootReposity.SubRepositories = append(rootReposity.SubRepositories, serviceReposity)
		var webReposity *Repository = &Repository {
			Name  : "web",
			Path  : pathRoot,
			RType : 2,
		}
		webReposity.ParentReposity = rootReposity
		rootReposity.SubRepositories = append(rootReposity.SubRepositories, webReposity)
		// rely sample
		var relySampleReposity *Repository = &Repository {
			Name  : "rely.sample",
			Path  : pathRoot + GetPathSeparator() + "rely",
			RType : 2,
		}
		relySampleReposity.ParentReposity = relyReposity
		// service sample
		var serviceSampleReposity *Repository = &Repository {
			Name  : "service.sample",
			Path  : pathRoot + GetPathSeparator() + "service",
			RType : 2,
		}
		serviceSampleReposity.ParentReposity = serviceReposity
		// web front
		var webFrontReposity *Repository = &Repository {
			Name  : "web.front",
			Path  : pathRoot + GetPathSeparator() + "web",
			RType : 2,
		}
		webFrontReposity.ParentReposity = webReposity
		// web rest
		var webRestReposity *Repository = &Repository {
			Name  : "web.rest",
			Path  : pathRoot + GetPathSeparator() + "web",
			RType : 2,
		}
		webRestReposity.ParentReposity = webReposity
		webReposity.SubRepositories = make([]*Repository, 2)
		webReposity.SubRepositories = append(webReposity.SubRepositories, webRestReposity)
		webReposity.SubRepositories = append(webReposity.SubRepositories, webFrontReposity)
		// create project
		rootReposity.Create()
	}
}
/*******************************************************************
*   template make
*/
type RootTmpStrcut struct {
	GroupName                   string
	ProjectName                 string
}

var RootTemplate string = `
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
	<modelVersion>4.0.0</modelVersion>
	<groupId>{{.GroupName}}</groupId>
	<artifactId>{{.ProjectName}}</artifactId>
	<version>0.0.1-SNAPSHOT</version>
	<packaging>pom</packaging>
	<modules>
		<module>rely</module>
		<module>web</module>
		<module>service</module>
	</modules>
	<build>
		<plugins>
			<plugin>
				<groupId>org.apache.maven.plugins</groupId>
				<artifactId>maven-compiler-plugin</artifactId>
				<version>3.1</version>
				<configuration>
					<source>1.8</source>
					<target>1.8</target>
					<encoding>UTF-8</encoding>
				</configuration>
			</plugin>
		</plugins>
	</build>
	<properties>
		<junit.version>4.5</junit.version>
		<druid.version>1.1.5</druid.version>
		<disruptor.version>3.3.7</disruptor.version>
		<mysql.version>5.1.44</mysql.version>
		<mybatis.version>3.4.5</mybatis.version>
		<mybatis.springboot.version>1.3.1</mybatis.springboot.version>
	</properties>
	<dependencyManagement>
		<dependencies>
			<!-- junit -->
			<dependency>
				<groupId>junit</groupId>
				<artifactId>junit</artifactId>
				<version>${junit.version}</version>
			</dependency>
			<!-- mybatis -->
			<dependency>
				<groupId>org.mybatis</groupId>
				<artifactId>mybatis</artifactId>
				<version>${mybatis.version}</version>
			</dependency>
			<dependency>
				<groupId>org.mybatis.spring.boot</groupId>
				<artifactId>mybatis-spring-boot-starter</artifactId>
				<version>${mybatis.springboot.version}</version>
			</dependency>
			<!-- mysql -->
			<dependency>
				<groupId>mysql</groupId>
				<artifactId>mysql-connector-java</artifactId>
				<version>${mysql.version}</version>
			</dependency>
			<!-- druid -->
			<dependency>
				<groupId>com.alibaba</groupId>
				<artifactId>druid</artifactId>
				<version>${druid.version}</version>
			</dependency>
			<!-- disruptor -->
			<dependency>
				<groupId>com.lmax</groupId>
				<artifactId>disruptor</artifactId>
				<version>${disruptor.version}</version>
			</dependency>
		</dependencies>
	</dependencyManagement>
	<dependencies>
		<dependency>
			<groupId>org.slf4j</groupId>
			<artifactId>slf4j-api</artifactId>
			<version>1.7.24</version>
		</dependency>
	</dependencies>
</project>
`

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