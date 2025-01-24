package dal

import (
	"github.com/xxxx2077/min_douyin_eCommerce/app/payment/biz/dal/mysql"
	"github.com/xxxx2077/min_douyin_eCommerce/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
