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

type User struct {
	ID           uint
	Name         string
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

	// 建表操作
	// db.AutoMigrate(&User{})
	// test_email := "666@163.com"
	// db.Create(&User{Name: "whyccc", Email: &test_email})
	// empty := "" // 这种方式也可以传递空值
	// db.Model(&User{ID: 1}).Updates(User{Email: &empty})

	/* 批量插入,最终会生成一个单一的sql语句 */
	// var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	// db.Create(&users)  // 一次全部提交
	//为什么不一次性提交所有的还要分批次，Sq1语句有长度限制
	// db.CreateInBatches(users, 2) // 2表示每次提交两条

	/* 通过map方式来创建 */
	// INSERT INTO `users` (`age`,`name`) VALUES (24,'ccc_map')
	db.Model(&User{}).Create(map[string]interface{}{
		"Name": "ccc_map", "Age": 24,
	})
}
