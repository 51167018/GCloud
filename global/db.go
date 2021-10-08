package global

import (
	"GCloud/models"
	"GCloud/pkg/conf"
	"fmt"
	"github.com/mattn/go-colorable"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var (
	DBEngine *gorm.DB
	err error
)

func InitDataBase() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DataBaseUserName,
		conf.DataBasePassWord,
		conf.DataBaseHost,
		conf.DataBasePort,
		conf.DataBaseName,
	)
	// newLogger 新建Logger
	var newLogger = logger.New(
		// io.writer同样使用colorable
		log.New(colorable.NewColorableStdout(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,       // 开启彩色打印
		},
	)
	DBEngine, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dns,
		DefaultStringSize: 256, // string 类型字段的默认长度
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //是否跳过外键约束
		// gorm日志模式：silent
		Logger:                 newLogger.LogMode(logger.Silent),
		SkipDefaultTransaction: true, //跳过默认事物
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gcloud_", //创建的表前缀
			SingularTable: false,   //复数创建表
		},
	})
	if err != nil {
		log.Println("连接数据库失败，请检查参数：", err)
		os.Exit(1)
	}
	_ = DBEngine.AutoMigrate(&models.User{},&models.File{},&models.Folder{})
	sqlDB, _ := DBEngine.DB()
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
