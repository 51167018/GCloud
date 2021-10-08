package conf

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	//Gin服务
	AppMode  string
	ServerPort string
	ServerHost string

	//数据库
	DataBaseDriver         string
	DataBaseHost     string
	DataBasePort     string
	DataBaseUserName     string
	DataBasePassWord string
	DataBaseName     string
)

func init()  {
	file, err := ini.Load("config/conf.ini")
	if err != nil {
		log.Println("配置文件加载失败")
	}
	//从本地加载数据库配置文件
	LoadDataBase(file)
	//从本地加载Gin配置文件
	LoadServer(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("Service").Key("AppMode").MustString("debug")
	ServerPort = file.Section("Service").Key("ServerPort").MustString(":3000")
	ServerHost = file.Section("Service").Key("ServerHost").MustString("127.0.0.1")
	log.Println("Gin配置文件加载成功")
}

func LoadDataBase(file *ini.File) {
	DataBaseDriver = file.Section("DataBase").Key("DataBaseDriver").MustString("debug")
	DataBaseHost = file.Section("DataBase").Key("DataBaseHost").MustString("localhost")
	DataBasePort = file.Section("DataBase").Key("DataBasePort").MustString("3306")
	DataBaseUserName = file.Section("DataBase").Key("DataBaseUserName").MustString("ginblog")
	DataBasePassWord = file.Section("DataBase").Key("DataBasePassWord").MustString("admin123")
	DataBaseName = file.Section("DataBase").Key("DataBaseName").MustString("ginblog")
	log.Println("数据库配置文件加载成功")
}
