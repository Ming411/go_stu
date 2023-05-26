package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 这里的User就是表名称
type User struct {
	ID           uint
	MyName       string  `gorm:"column:name"`
	Email        *string // 使用指针的方法也可以解决 空值问题
	Age          uint8   // uint8 只能表示 0 到 255 之间的整数
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
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
	/* 方式一 */
	db.First(&user)
	user.MyName = "whyccc2"
	user.Age = 66
	// Save方法update和create操作
	db.Save(&user)

	/* 方式二 */
	db.Model(&User{}).Where("name = ?", "hello").Update("age", 18)
}
