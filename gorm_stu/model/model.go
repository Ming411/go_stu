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

type Product struct {
	gorm.Model // Model中包含一些基础属性 id CreatedAt UpdatedAt DeletedAt
	// Code       string
	Code  sql.NullString
	Price uint // 正整型
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
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: sql.NullString{String: "D42", Valid: true}, Price: 100})

	// Read
	var product Product
	// db.First(&product, 1)                 // 根据整型主键查找
	// db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	// db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	// 非零值字段 即 非空串 0 等
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// 如果想设置空值的话,需要将类型设置为 sql.Nullxxxxx
	db.Model(&product).Updates(Product{Price: 200, Code: sql.NullString{String: "", Valid: true}})

	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product 逻辑删除，并非物理删除
	// db.Delete(&product, 1)
}
