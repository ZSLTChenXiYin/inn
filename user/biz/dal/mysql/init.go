package mysql

import (
	"os"
	"user/biz/dal/model"
	"user/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	if os.Getenv("GIN_ENV") != "online" {
		err = DB.AutoMigrate(&model.User{})
		if err != nil {
			panic(err)
		}
	}
}
