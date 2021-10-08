package user

import (
	"GCloud/dao"
	"GCloud/models"
	"GCloud/pkg/errcode"
	"GCloud/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

var (
	code int
)

func Register(c *gin.Context) {
	//创建一个对象来绑定Json
	var User models.User
	err := c.ShouldBindJSON(&User)
	if err != nil {
		log.Println("json绑定失败")
		return
	}
	code = dao.CreateUser(&User)
	util.NewResponse(c).ToResponse(code, nil)
}

//LoginUser 登陆
func LoginUser(c *gin.Context)  {
	username := c.PostForm("username")
	password := []byte(c.PostForm("password"))
	
	passwordForDataBase, code := dao.FindUserNameAndPassWord(username)
	if code==errcode.SELECT_DATABASE_ERROR{
		util.NewResponse(c).ToResponse(code, nil)
		return
	}
	ok := util.ValidatePasswords(passwordForDataBase, password)
	if ok {
		util.NewResponse(c).ToResponse(errcode.SUCCESS, nil)
		return
	}
	util.NewResponse(c).ToResponse(errcode.ERROR_PASSWORD_ERROR, nil)
}

//AllUserInfo 查询所有用户信息
func AllUserInfo(c *gin.Context)  {
	users,code := dao.FindAllUser()
	util.NewResponse(c).ToResponse(code, users)
}

//FindUserInfoByColumn 查询单个用户
func FindUserInfoByColumn(c *gin.Context)  {
	var u models.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		log.Println("json绑定失败")
		return
	}
	user, code := dao.FindUserInfoByColumn(&u)
	util.NewResponse(c).ToResponse(code, user)
}

func UpdateUserInfo(c *gin.Context)  {
	var u models.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("参数转换失败")
		return
	}
	err = c.ShouldBindJSON(&u)
	if err != nil {
		log.Println("json绑定失败")
		return
	}
	code := dao.UpdateUserInfo(id,&u)
	util.NewResponse(c).ToResponse(code, nil)
}

//DeleteUser 删除用户
func DeleteUser(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("参数转换失败")
		return
	}
	code := dao.DeleteUser(id)
	util.NewResponse(c).ToResponse(code, nil)
}

//FindUserFileById 查询该用户的所有文件
func FindUserFileById(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("参数转换失败")
		return
	}
	userFiles, code := dao.FindUserFileById(id)
	util.NewResponse(c).ToResponse(code, userFiles)
}

//FindUserFileAndFolderById 查询该用户的所有文件
func FindUserFileAndFolderById(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("参数转换失败")
		return
	}
	userFilesAndFolder, code := dao.FindUserFileAndFolderById(id)
	util.NewResponse(c).ToResponse(code, userFilesAndFolder)
}

