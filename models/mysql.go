package models

import (
	"fmt"
    "os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	SqlSession, err = gorm.Open(DRIVER, dsn)
	if err != nil {
		panic(err)
	}
	SqlSession.AutoMigrate(&User{}, &Video{})
	return SqlSession.DB().Ping()
}

func CloseMySQL() {
	err := SqlSession.Close()
	if err != nil {
		panic(err)
	}
}
