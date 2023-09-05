package models

import (
	"fmt"
    "os"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
	//"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

const DRIVER = "mysql"

var SqlSession *gorm.DB

func InitMySql(conf MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		// conf.UserName,
		// conf.Password,
		// conf.Url,
		// conf.Port,
		conf.DBName,
	)
	SqlSession, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

    // 创建表格
	SqlSession.AutoMigrate(&User{}, &Video{})
    // SqlSession.Exec("alter table users add  CONSTRAINT pk_users primary key (user_id) ;")
	return nil
}

func CloseMySQL() {
	db, _ := SqlSession.DB()
    err := db.Close()
	if err != nil {
		panic(err)
	}
}
