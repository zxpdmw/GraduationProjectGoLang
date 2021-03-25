package env

import (
	"os"
	"runtime"
)

var ostype = runtime.GOOS

func GetProjectPath() string {
	var projectPath string
	projectPath, _ = os.Getwd()
	return projectPath
}

func GetIp() string {
	ip := "http://"
	if ostype == "windows" {
		ip += "localhost:8080"
		ip += "/swagger/doc.json"
	} else if ostype == "linux" {
		ip += "39.96.113.190:8080"
		ip += "/swagger/doc.json"
	}
	return ip
}

func GetTemplatePath() string {
	path := GetProjectPath()
	if ostype == "windows" {
		path = path + "\\" + "templates\\*"
	} else if ostype == "linux" {
		path = path + "/" + "templates/*"
	}
	return path
}

func GetConfigPath() string {
	path := GetProjectPath()
	if ostype == "windows" {
		path = path + "\\" + "config\\"
	} else if ostype == "linux" {
		path = path + "/" + "config/"
	}
	return path
}

func GetConLogPath() string {
	path := GetProjectPath()
	if ostype == "windows" {
		path = path + "\\log\\"
	} else if ostype == "linux" {
		path = path + "/log/"
	}
	return path
}
