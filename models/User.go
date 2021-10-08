package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model 
	UserName string   `json:"username"`//用户名
	PassWord string   `json:"password"`//用户密码
	Email    string   `json:"email" gorm:"type:varchar(100);unique_index"`
	Avatar   string   `json:"avatar"`//用户头像
	Status   int      `json:"status"`//用户的状态 0正常 1封禁
	Role     int      `json:"role"`//用户角色,0管理员，1普通用户
	Files    []File   `json:"files"`//用户所属的文件
	Folders  []Folder `json:"folders"`//用户所属的文件夹
}

