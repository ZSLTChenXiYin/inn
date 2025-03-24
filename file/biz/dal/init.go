package dal

import (
	"file/biz/dal/mysql"
	"file/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
