package main

import (
	"database/sql"
	"errors"
	"fmt"
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
	var users User
	// 获取第一条记录（主键升序）
	db.First(&users)
	// SELECT * FROM users ORDER BY id LIMIT 1;
	// fmt.Println(users.ID, "First") // 它会将第一次查询到的数据挂载到这个变量上
	// 获取一条记录，没有指定排序字段
	db.Take(&users)
	// SELECT * FROM users LIMIT 1;

	// 获取最后一条记录（主键降序）
	db.Last(&users)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	result := db.First(&users)
	// result.RowsAffected // 返回找到的记录数
	// result.Error        // returns error or nil

	// 检查 ErrRecordNotFound 是否出现错误
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("查找数据失败")
	}

	// 通过主键查询
	db.First(&users, 4)              // 内部会自动做数据类型转换
	db.First(&users, []int{1, 2, 3}) // 查找主键为 1 or 2 or 3的数据，默认只输出第一条

	var users2 []User // 检索全部对象
	ret := db.Find(&users2)
	fmt.Println("总条数", ret.RowsAffected)
	for _, user := range users2 {
		fmt.Println(user)
	}
}
