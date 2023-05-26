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

type User struct {
	gorm.Model
	// 内部会多给我建立一张关系表user_languages
	Languages []Language `gorm:"many2many:user_languages;"`
}
type Language struct {
	gorm.Model
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

	// db.AutoMigrate(&User{})

	// languages := []Language{
	// 	{Name: "py"}, {
	// 		Name: "c#",
	// 	},
	// }
	// // languages := []Language{}
	// // languages = append(languages, Language{Name: "go"})
	// // languages = append(languages, Language{Name: "ts"})
	// user := User{
	// 	Languages: languages,
	// }
	// db.Create(&user)

	// var user User
	// // db.Preload("Languages").First(&user)
	// db.Preload("Languages").Where("ID = ?", 2).First(&user)
	// for _, language := range user.Languages {
	// 	fmt.Println(language.Name)
	// }

	var user User
	db.First(&user)
	var languages []Language
	// 通过已知用户信息查找
	_ = db.Model(&user).Association("Languages").Find(&languages)
	for _, language := range languages {
		fmt.Println(language.Name)
	}
}
