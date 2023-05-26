package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Company struct {
	gorm.Model
	Name    string
	AddTime time.Time // 每次记录创建时自动更新事件，搭配钩子函数
}

func (c *Company) BeforeCreate(tx *gorm.DB) (err error) {
	c.AddTime = time.Now()
	return
}

// 在gorm中可以通过给某一个struct,添加TableName.方法来自定义表名
// func (Company) TableName() string {
// 	return "my_company"
// }

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
		// 这个配置不能与tablename同时设置
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "mxshop_", // 所有表都会加上这个前缀
		},
		Logger: newLogger,
	})
	if err != nil {
		panic("数据库连接失败")
	}

	db.AutoMigrate(&Company{})
}
