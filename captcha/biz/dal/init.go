package dal

import (
	"captcha/biz/dal/redis"
)

func Init() {
	redis.Init()
	// mysql.Init()
}
