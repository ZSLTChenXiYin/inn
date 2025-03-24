package mysql

import (
	"file/biz/dal/model"
	"file/conf"
	"os"

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
		err = DB.AutoMigrate(&model.File{})
		if err != nil {
			panic(err)
		}
	}
}
