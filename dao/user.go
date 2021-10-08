package dao

import (
	"GCloud/global"
	"GCloud/models"
	"GCloud/pkg/errcode"
	"GCloud/pkg/util"
	"gorm.io/gorm"
	"log"
)

//CreateUser 创建用户
func CreateUser(user *models.User) int {
	existUserName := CheckUserNameExist(user)
	if existUserName == 0 {
		user.PassWord = util.HashAndSalt([]byte(user.PassWord))
		if err := global.DBEngine.Model(user).Create(user).Error; err != nil {
			log.Println("用户创建失败-", err.Error())
			return errcode.INSERT_DATABASE_ERROR
		}
		return errcode.SUCCESS
	}
	return errcode.ERROR_USERNAME_EXIST
}


//CheckUserNameExist 检查用户名是否存在
func CheckUserNameExist(user *models.User) int64 {
	var total int64
	if err := global.DBEngine.Debug().Model(user).Where("user_name = ?", user.UserName).Find(user).Count(&total).Error; err != nil {
		log.Println("检测用户名失败-", err.Error())
		return 1
	}
	return total
}

//FindAllUser 查询所有用户
func FindAllUser() ([]models.User, int) {
	var users []models.User
	if err := global.DBEngine.Debug().Model(&users).Find(&users).Error; err != nil {
		log.Println("查询所有用户失败-", err.Error())
		return users, errcode.SELECT_DATABASE_ERROR
	}
	return users, errcode.SUCCESS
}

//FindUserNameAndPassWord 查询用户账户密码
func FindUserNameAndPassWord(username string) (string, int) {
	var password string
	if err := global.DBEngine.Debug().Model(&models.User{}).Select("pass_word").Where("user_name = ?",username).First(&password).Error; err != nil {
		log.Println("查询用户密码失败-", err.Error())
		return password, errcode.SELECT_DATABASE_ERROR
	}
	return password, errcode.SUCCESS
}

//FindUserInfoByColumn 按条件查询单个用户
func FindUserInfoByColumn(u *models.User) (models.User, int) {
	var user models.User
	if err := global.DBEngine.Debug().Model(&user).Where(u).First(&user).Error; err != nil {
		log.Println("查询单个用户失败-", err.Error())
		return user, errcode.SELECT_DATABASE_ERROR
	}
	return user, errcode.SUCCESS
}

//UpdateUserInfo 修改用户信息
func UpdateUserInfo(id int,u *models.User)int {
	if err := global.DBEngine.Debug().Model(&models.User{}).Where("id = ?",id).Updates(u).Error; err != nil {
		log.Println("更新用户失败-", err.Error())
		return errcode.UPDATE_DATABASE_ERROR
	}
	return errcode.SUCCESS
}

//DeleteUser 删除用户
func DeleteUser(id int) int {
	if err := global.DBEngine.Debug().Model(&models.User{}).Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		log.Println("删除用户失败-", err.Error())
		return errcode.DELETE_DATABASE_ERROR
	}
	return errcode.SUCCESS
}

//FindUserFileById 查询该用户的所有文件
func FindUserFileById(id int) (models.User, int) {
	var User models.User
	err := global.DBEngine.Debug().Model(&models.User{}).Preload("Files", func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", id)
	}).First(&User).Error
	if err != nil {
		log.Println("查询该用户的所有文件失败-", err.Error())
		return User, errcode.SELECT_DATABASE_ERROR
	}
	return User, errcode.SUCCESS
}

//FindUserFileAndFolderById 查询该用户下的所有文件夹和文件
func FindUserFileAndFolderById(id int) (models.User, int) {
	var User models.User
	err := global.DBEngine.Debug().Model(&models.User{}).Preload("Files", func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", id)
	}).Preload("Folders", func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", id)
	}).First(&User).Error
	if err != nil {
		log.Println("查询该用户的所有文件和文件夹失败-", err.Error())
		return User, errcode.SELECT_DATABASE_ERROR
	}
	return User, errcode.SUCCESS
}
