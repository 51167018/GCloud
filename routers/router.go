package routers

import (
	"GCloud/pkg/conf"
	"GCloud/routers/api/v1/file"
	"GCloud/routers/api/v1/user"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"log"
)

func InitRouter()  {
	// 启用gin的日志输出带颜色
	gin.ForceConsoleColor()
	// 替换默认Writer（关键步骤）
	gin.DefaultWriter = colorable.NewColorableStdout()
	//设置启动模式
	gin.SetMode(conf.AppMode)
	//创建Gin路由
	r := gin.Default()
	//中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//跨域
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	})
	//创建路由组
	allRouterGroup := r.Group("api/v1")
	userGroup := allRouterGroup.Group("user")
	//用户组路由
	{
		userGroup.POST("/Register",user.Register)
		userGroup.GET("/UserInfoList",user.AllUserInfo)
		userGroup.GET("/UserInfo",user.FindUserInfoByColumn)
		userGroup.DELETE("/DeleteUser/:id",user.DeleteUser)
		userGroup.PUT("/UserInfo/:id",user.UpdateUserInfo)
		userGroup.POST("/UserLogin",user.LoginUser)

		userGroup.GET("/UserFile/:id",user.FindUserFileById)
		userGroup.GET("/UserFileAndFolder/:id",user.FindUserFileAndFolderById)
	}
	fileGroup := allRouterGroup.Group("file")
	{
		fileGroup.GET("/checkChunk",file.CheckChunk)
		fileGroup.GET("/megerChunk",file.MegerChunk)
		fileGroup.POST("/uploadChunk",file.UploadChunk)
	}
	err := r.Run(conf.ServerPort)
	if err != nil {
		log.Println("Gin路由启动失败")
	}
}
