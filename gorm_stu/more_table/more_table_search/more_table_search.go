package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// User属于Company CompanyID 是外键
type User struct {
	gorm.Model
	Name      string
	CompanyID int     // 名字会自动转为蛇形company_id
	Company   Company // 这个的作用是为了方便执行查询
}

type Company struct {
	ID   int
	Name string
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(172.20.57.237:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	// 配置日志输出  可查看具体执行的SQL
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("数据库连接失败")
	}
	var user User
	// db.First(&user)
	// 关联查询
	/* 方式一  执行两步SQL先查询company然后在查询user */
	// db.Preload("Company").First(&user)
	/* 方式二 通过join一步获取*/
	db.Joins("Company").First(&user)

	fmt.Println(user)
}
