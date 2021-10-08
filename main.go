package main

import (
	"GCloud/global"
	"GCloud/routers"
)

func main() {
	//初始化数据库
	global.InitDataBase()
	//初始化Gin路由
	routers.InitRouter()
}
