package main

import(
	"os"
	"text/template"
	"strings"
)

type MavenConfig struct {
	PjConfig        ProjectConfig
	Group           string   `group id`
	PKName          string   `package name`
}

/*******************************************************************
*   folder make
*/
func (this *MavenConfig) createRootFile(path string, file string) (string, error){
	var rootData RootTmpStrcut = RootTmpStrcut{
		GroupName     : this.Group,
		ProjectName   : this.PjConfig.Name,
	}
	tmpl, err:= template.New("Root").Parse(RootTemplate)
	if err != nil {
		return "", err
	}
	var filePath string = path + GetPathSeparator() + file
	f, er := os.Create(filePath)
	if er != nil {
		return "", er
	}
	err = tmpl.Execute(f, rootData)
	if err != nil {
		return "", err	
	}
	return "", nil
}

func (this *MavenConfig) createWebFile(path string, file string) (string, error){
	var webData WebTmpStrcut = WebTmpStrcut{
		ParentTmp  :  RootTmpStrcut{
			GroupName     : this.Group,
			ProjectName   : this.PjConfig.Name,
		},
	}
	tmpl, err:= template.New("Web").Parse(WebTemplate)
	if err != nil {
		return "", err
	}
	var filePath string = path + GetPathSeparator() + file
	f, er := os.Create(filePath)
	if er != nil {
		return "", er
	}
	err = tmpl.Execute(f, webData)
	if err != nil {
		return "", err	
	}
	return "", nil
}

func (this *MavenConfig) createServiceFile(path string, file string) (string, error){
	var serviceData ServiceTmpStrcut = ServiceTmpStrcut{
		ParentTmp  :  RootTmpStrcut{
			GroupName     : this.Group,
			ProjectName   : this.PjConfig.Name,
		},
	}
	tmpl, err:= template.New("service").Parse(ServiceTemplate)
	if err != nil {
		return "", err
	}
	var filePath string = path + GetPathSeparator() + file
	f, er := os.Create(filePath)
	if er != nil {
		return "", er
	}
	err = tmpl.Execute(f, serviceData)
	if err != nil {
		return "", err	
	}
	return "", nil
}

func (this *MavenConfig) createServiceSamplePom(path string, file string) (string, error){
	var serviceSampleData ServiceSampleStrcut = ServiceSampleStrcut{
		ParentTmp  :  RootTmpStrcut{
			GroupName     : this.Group,
			ProjectName   : this.PjConfig.Name,
		},
	}
	tmpl, err:= template.New("service").Parse(ServiceSampleTemplate)
	if err != nil {
		return "", err
	}
	var filePath string = path + GetPathSeparator() + file
	f, er := os.Create(filePath)
	if er != nil {
		return "", er
	}
	err = tmpl.Execute(f, serviceSampleData)
	if err != nil {
		return "", err	
	}
	return "", nil
}

func (this *MavenConfig) createServiceSampleJava(path string, file string) (string, error){
	return "", nil
}

func (this *MavenConfig) createRelyFile(path string, file string) (string, error){
	var relyData RelyTmpStrcut = RelyTmpStrcut{
		ParentTmp  :  RootTmpStrcut{
			GroupName     : this.Group,
			ProjectName   : this.PjConfig.Name,
		},
	}
	tmpl, err:= template.New("Rely").Parse(RelyTemplate)
	if err != nil {
		return "", err
	}
	var filePath string = path + GetPathSeparator() + file
	f, er := os.Create(filePath)
	if er != nil {
		return "", er
	}
	err = tmpl.Execute(f, relyData)
	if err != nil {
		return "", err	
	}
	return "", nil
}

func (this *MavenConfig) createRelySampleJavaFile(path string, file string) (string, error){
	var relySampleJavaData RelySampleJavaStruct = RelySampleJavaStruct{
		ParentTmp	:  RootTmpStrcut{
			GroupName     : this.Group,
			ProjectName   : this.PjConfig.Name,
		},
		PackageName       : this.PKName,
	}
	tmpl, err := template.New("RelySampleJava").Parse(RelySampleJavaTemplate)
	if err != nil {
		return "", err
	}
	var filePath string = path + GetPathSeparator() + file
	f, er := os.Create(filePath)
	if er != nil {
		return "", er
	}
	err = tmpl.Execute(f, relySampleJavaData)
	if err != nil {
		return "", err	
	}
	return "", nil
}

func (this *MavenConfig) createRelySampleFile(path string, file string) (string, error){
	var relyData RelySampleStruct = RelySampleStruct{
		ParentTmp	:  RootTmpStrcut{
			GroupName     : this.Group,
			ProjectName   : this.PjConfig.Name,
		},
	}
	tmpl, err := template.New("RelySample").Parse(RelySampleTemplate)
	if err != nil {
		return "", err
	}
	var filePath string = path + GetPathSeparator() + file
	f, er := os.Create(filePath)
	if er != nil {
		return "", er
	}
	err = tmpl.Execute(f, relyData)
	if err != nil {
		return "", err	
	}
	return "", nil
}


func (this *MavenConfig) Making() {
	projectName   := this.PjConfig.Name
	//groupName     := this.Group
	//pakcageName   := this.PKName
	path          := this.PjConfig.Path
	// set root folder
	var pathRoot string = path + GetPathSeparator() + projectName
	if !IsExist(pathRoot) {
		var rootReposity *Repository = &Repository{
			Name       : projectName,
			Path       : path,
			RType      : 2,
			CreateFold : CreateFolder,
		}
		rootReposity.ParentReposity = nil
		rootReposity.SubRepositories = make([]*Repository, 4)
		///////////////////////////////////////////////////////////////////////////////////////
		// root pom define begin
		var rootPomReposity *Repository = &Repository{
			Name       : "pom.xml",
			Path       : pathRoot,
			RType      : 1,
			CreateFile : this.createRootFile,
		}
		rootPomReposity.ParentReposity = rootReposity
		rootReposity.SubRepositories = append(rootReposity.SubRepositories, rootPomReposity)	
		///////////////////////////////////////////////////////////////////////////////////////
		// rely define begin
		//		
		var relyReposity *Repository = &Repository {
			Name       : "rely",
			Path       : pathRoot,
			RType      : 2,
			CreateFold : CreateFolder,
		}
		relyReposity.ParentReposity = rootReposity
		relyReposity.SubRepositories = make([]*Repository, 2)
		rootReposity.SubRepositories = append(rootReposity.SubRepositories, relyReposity)
		var relyPomReposity *Repository = &Repository{
			Name       : "pom.xml",
			Path       : pathRoot + GetPathSeparator() + "rely",
			RType      : 1,
			CreateFile : this.createRelyFile,
		}
		relyPomReposity.ParentReposity = relyReposity
		relyReposity.SubRepositories = append(relyReposity.SubRepositories, relyPomReposity)
		///////////////////////////////////////////////////////////////////////////////////////
		// serivce define begin
		//
		var serviceReposity *Repository = &Repository {
			Name       : "service",
			Path       : pathRoot,
			RType      : 2,
			CreateFold : CreateFolder,
		}
		serviceReposity.ParentReposity = rootReposity
		serviceReposity.SubRepositories = make([]*Repository, 2)
		rootReposity.SubRepositories = append(rootReposity.SubRepositories, serviceReposity)
		var serivcePomReposity *Repository = &Repository{
			Name       : "pom.xml",
			Path       : pathRoot + GetPathSeparator() + "service",
			RType      : 1,
			CreateFile : this.createServiceFile,
		}
		serivcePomReposity.ParentReposity = serviceReposity
		serviceReposity.SubRepositories = append(serviceReposity.SubRepositories, serivcePomReposity)

		var serviceSampleReposity *Repository = &Repository {
			Name       : "service.sample",
			Path       : pathRoot + GetPathSeparator() + "service",
			RType      : 2,
			CreateFold : CreateFolder,
		}
		serviceSampleReposity.ParentReposity = serviceReposity
		serviceSampleReposity.SubRepositories = make([]*Repository, 1)
		serviceReposity.SubRepositories = append(serviceReposity.SubRepositories, serviceSampleReposity)

		var serviceSamplePomReposity *Repository = &Repository {
			Name       : "pom.xml",
			Path       : pathRoot + GetPathSeparator() + "service" + GetPathSeparator() + "service.sample",
			RType      : 1,
			CreateFile : this.createServiceSamplePom,
		}
		serviceSamplePomReposity.ParentReposity = serviceSampleReposity
		serviceSamplePomReposity.SubRepositories = make([]*Repository, 0)
		serviceSampleReposity.SubRepositories = append(serviceSampleReposity.SubRepositories, serviceSamplePomReposity)


		var serviceSampleSrcReposity *Repository = &Repository {
			Name       : "src",
			Path       : pathRoot + GetPathSeparator() + "service" + GetPathSeparator() + "service.sample",
			RType      : 2,
			CreateFold : CreateFolder,
		}
		serviceSampleSrcReposity.ParentReposity = serviceSampleReposity
		serviceSampleSrcReposity.SubRepositories = make([]*Repository, 2)
		serviceSampleReposity.SubRepositories = append(serviceSampleReposity.SubRepositories, serviceSampleSrcReposity)

		var servicePK string =  ""
		if this.PKName != "" {
			servicePK = "main.java." + this.PKName + ".sample.service.impl"
		} else {
			servicePK = "main.java.sample.service.impl"
		}
		var serviceParentReposity *Repository = serviceSampleSrcReposity
		var serviceSampleSrcPath string = serviceSampleSrcReposity.Path + GetPathSeparator() + "src"
		var svcTmpFolder *Repository = nil
		serviceFS := strings.SplitN(servicePK, ".", -1)
		for _, f := range serviceFS {
			svcTmpFolder = &Repository {
				Name       : f,
				Path       : serviceSampleSrcPath,
				RType      : 2,
				CreateFold : CreateFolder,
			}
			svcTmpFolder.ParentReposity = serviceParentReposity
			svcTmpFolder.SubRepositories = make([]*Repository, 1)
			serviceParentReposity.SubRepositories = append(serviceParentReposity.SubRepositories, svcTmpFolder)
			serviceParentReposity = svcTmpFolder
			serviceSampleSrcPath = serviceSampleSrcPath + GetPathSeparator() + f
		}
		var serviceSampleJavaCode *Repository = &Repository {
			Name       : "ServiceSampleImpl.java",
			Path       : svcTmpFolder.Path + GetPathSeparator() + svcTmpFolder.Name,
			RType      : 1,
			CreateFile : this.createServiceSampleJava,			
		}
		serviceSampleJavaCode.ParentReposity = serviceParentReposity
		serviceParentReposity.SubRepositories = append(serviceParentReposity.SubRepositories, serviceSampleJavaCode)

		///////////////////////////////////////////////////////////////////////////////////////
		// serivce define end
		//
		var webReposity *Repository = &Repository {
			Name       : "web",
			Path       : pathRoot,
			RType      : 2,
			CreateFold : CreateFolder,
		}
		webReposity.ParentReposity = rootReposity
		rootReposity.SubRepositories = append(rootReposity.SubRepositories, webReposity)
		///////////////////////////////////////////////////////////////////////////////////////
		// rely sample define begin
		//
		var relySampleReposity *Repository = &Repository {
			Name       : "rely.sample",
			Path       : pathRoot + GetPathSeparator() + "rely",
			RType      : 2,
			CreateFold : CreateFolder,
		}
		relySampleReposity.ParentReposity = relyReposity
		relySampleReposity.SubRepositories = make([]*Repository, 0)
		relyReposity.SubRepositories = append(relyReposity.SubRepositories, relySampleReposity)
		var relySamplePomReposity *Repository = &Repository{
			Name       : "pom.xml",
			Path       : pathRoot + GetPathSeparator() + "rely" + GetPathSeparator() + "rely.sample",
			RType      : 1,
			CreateFile : this.createRelySampleFile,
		}
		relySamplePomReposity.ParentReposity = relySampleReposity
		relySampleReposity.SubRepositories = append(relySampleReposity.SubRepositories, relySamplePomReposity)
		var relySampleSrcReposity *Repository = &Repository {
			Name  : "src",
			Path  : pathRoot + GetPathSeparator() + "rely" + GetPathSeparator() + "rely.sample",
			RType : 2,
			CreateFold : CreateFolder,
		}
		relySampleSrcReposity.ParentReposity = relySampleReposity
		relySampleReposity.SubRepositories = append(relySampleReposity.SubRepositories, relySampleSrcReposity)
		relySampleSrcReposity.SubRepositories = make([]*Repository, 0)
		var relySampleMainReposity *Repository = &Repository {
			Name  : "main",
			Path  : pathRoot + GetPathSeparator() + "rely" + GetPathSeparator() + "rely.sample" + GetPathSeparator() + "src",
			RType : 2,
		}
		relySampleMainReposity.ParentReposity = relySampleReposity
		relySampleMainReposity.SubRepositories = make([]*Repository, 0)
		relySampleSrcReposity.SubRepositories = append(relySampleSrcReposity.SubRepositories, relySampleMainReposity)
		var relySampleJavaReposity *Repository = &Repository {
			Name       : "java",
			Path       : pathRoot + GetPathSeparator() + "rely" + GetPathSeparator() + "rely.sample" + GetPathSeparator() + "src" + GetPathSeparator() + "main",
			RType      : 2,
			CreateFold : CreateFolder,
		}
		relySampleJavaReposity.ParentReposity = relySampleMainReposity
		relySampleJavaReposity.SubRepositories = make([]*Repository, 0)
		relySampleMainReposity.SubRepositories = append(relySampleMainReposity.SubRepositories, relySampleJavaReposity)
		// make package & sample java file
		var parentReposity *Repository = relySampleJavaReposity
		if this.PKName != "" {
			var relyPackageName = this.PKName + ".sample.service"
			var p string = pathRoot + GetPathSeparator() + "rely" + GetPathSeparator() + "rely.sample" + GetPathSeparator() + "src" + GetPathSeparator() + "main" + GetPathSeparator() + "java"
			var tmpFolder *Repository = nil
			fs := strings.SplitN(relyPackageName, ".", -1)
			for _, f := range fs {
				tmpFolder = &Repository {
					Name       : f,
					Path       : p,
					RType      : 2,
					CreateFold : CreateFolder,
				}
				tmpFolder.ParentReposity = parentReposity
				tmpFolder.SubRepositories = make([]*Repository, 1)
				parentReposity.SubRepositories = append(parentReposity.SubRepositories, tmpFolder)
				parentReposity = tmpFolder
				p = p + GetPathSeparator() + f
			}
			var relySampleJavaReposity *Repository = &Repository{
				Name       : "RelySample.java",
				Path       : tmpFolder.Path + GetPathSeparator() + tmpFolder.Name,
				RType      : 1,
				CreateFile : this.createRelySampleJavaFile,				
			}
			relySampleJavaReposity.ParentReposity = tmpFolder
			tmpFolder.SubRepositories = append(tmpFolder.SubRepositories, relySampleJavaReposity)
		}
		///////////////////////////////////////////////////////////////////////////////////////
		// rely sample define end

		///////////////////////////////////////////////////////////////////////////////////////
		// web define begin
		//
		// web front
		var webFrontReposity *Repository = &Repository {
			Name       : "web.front",
			Path       : pathRoot + GetPathSeparator() + "web",
			RType      : 2,
			CreateFold : CreateFolder,
		}
		webFrontReposity.ParentReposity = webReposity
		// web rest
		var webRestReposity *Repository = &Repository {
			Name       : "web.rest",
			Path       : pathRoot + GetPathSeparator() + "web",
			RType      : 2,
			CreateFold : CreateFolder,
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
/**
*  rely
*/
type RelyTmpStrcut struct {
	ParentTmp                    RootTmpStrcut
}
var RelyTemplate string = `
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
	<modelVersion>4.0.0</modelVersion>
	<parent>
		<groupId>{{.ParentTmp.GroupName}}</groupId>
		<artifactId>{{.ParentTmp.ProjectName}}</artifactId>
		<version>0.0.1-SNAPSHOT</version>
	</parent>
	<artifactId>{{.ParentTmp.ProjectName}}.rely</artifactId>
	<packaging>pom</packaging>
	<modules>
        <module>rely.sample</module>
	</modules>
</project>
`
type RelySampleStruct struct {
	ParentTmp                RootTmpStrcut
}
var RelySampleTemplate string = `
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
	<modelVersion>4.0.0</modelVersion>
	<parent>
		<groupId>{{.ParentTmp.GroupName}}</groupId>
		<artifactId>{{.ParentTmp.ProjectName}}.rely</artifactId>
		<version>0.0.1-SNAPSHOT</version>
	</parent>
	<artifactId>{{.ParentTmp.ProjectName}}.rely.sample</artifactId>
	<packaging>jar</packaging>
</project>
`
type RelySampleJavaStruct struct {
	ParentTmp					RootTmpStrcut
	PackageName                 string
}
var RelySampleJavaTemplate string = `
package {{.PackageName}}.sample.service;

public interface RelySample {
	string sample();
}
`
/**
*  service
*/
type ServiceTmpStrcut struct {
	ParentTmp                    RootTmpStrcut
}
var ServiceTemplate string = `
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
	<modelVersion>4.0.0</modelVersion>
	<parent>
		<groupId>{{.ParentTmp.GroupName}}</groupId>
		<artifactId>{{.ParentTmp.ProjectName}}</artifactId>
		<version>0.0.1-SNAPSHOT</version>
	</parent>
	<artifactId>{{.ParentTmp.ProjectName}}.service</artifactId>
	<packaging>pom</packaging>
	<modules>
        <module>service.sample</module>
  	</modules>		
</project>
`
type ServiceSampleStrcut struct {
	ParentTmp                    RootTmpStrcut
}
var ServiceSampleTemplate string = `
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
	<modelVersion>4.0.0</modelVersion>
	<parent>
		<groupId>{{.ParentTmp.GroupName}}</groupId>
		<artifactId>{{.ParentTmp.ProjectName}}.service</artifactId>
		<version>0.0.1-SNAPSHOT</version>
	</parent>
	<artifactId>{{.ParentTmp.ProjectName}}.service.sample</artifactId>
	<packaging>jar</packaging>	
</project>
`
var ServiceSampleJavaTemplate string = 
`package {{.PackageName}}.sample.service.impl;
import {{.PackageName}}.sample.serivce;
public class ServiceSampleImpl implements RelySample {
	@Override
	public String sample() {
		return "hello world!";
	}
}
`

/**
*  web
*/
type WebTmpStrcut struct {
	ParentTmp                    RootTmpStrcut
}
var WebTemplate string = `
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
	<modelVersion>4.0.0</modelVersion>
	<parent>
		<groupId>{{.ParentTmp.GroupName}}</groupId>
		<artifactId>{{.ParentTmp.ProjectName}}</artifactId>
		<version>0.0.1-SNAPSHOT</version>
	</parent>
	<artifactId>{{.ParentTmp.ProjectName}}.web</artifactId>
	<packaging>pom</packaging>
	<modules>
        <module>web.rest</module>
  	</modules>
	<properties>
		<spring.boot.version>1.5.2.RELEASE</spring.boot.version>
	</properties>
	<build>
		<plugins>
			<plugin>
				<groupId>org.springframework.boot</groupId>
				<artifactId>spring-boot-maven-plugin</artifactId>
				<version>${spring.boot.version}</version>
				<executions>
					<execution>
						<goals>
							<goal>repackage</goal>
						</goals>
					</execution>
				</executions>
			</plugin>
		</plugins>
	</build>
	<dependencyManagement>
		<!-- web project using the spring boot -->
		<dependencies>
			<!-- import spring boot basic dependency -->
			<dependency>
				<groupId>org.springframework.boot</groupId>
				<artifactId>spring-boot-starter-web</artifactId>
				<type>pom</type>
		        <scope>import</scope>
				<version>${spring.boot.version}</version>
			</dependency>
		</dependencies>		
	</dependencyManagement>
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