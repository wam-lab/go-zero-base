package svc

import (
	"github.com/tal-tech/go-queue/dq"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/yguilai/timetable-micro/services/captcha/rpc/captchaclient"
	"github.com/yguilai/timetable-micro/services/jwt/rpc/jwtclient"
	"github.com/yguilai/timetable-micro/services/user/model"
	"github.com/yguilai/timetable-micro/services/user/rpc/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Conf       config.Config
	Db         *gorm.DB
	Redis      *redis.Redis
	Producer   dq.Producer
	JwtRpc     jwtclient.Jwt
	CaptchaRpc captchaclient.Captcha
}

func NewServiceContext(c config.Config) *ServiceContext {
	// new gorm and auto migrate
	db, err := newGorm(c.DataSource)
	if err != nil {
		panic(err)
	}

	err = model.AutoMigrate(db)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Conf:       c,
		Db:         db,
		Redis:      c.RedisConf.NewRedis(),
		Producer:   dq.NewProducer(c.Beanstalks),
		JwtRpc:     jwtclient.NewJwt(zrpc.MustNewClient(c.JwtRpc)),
		CaptchaRpc: captchaclient.NewCaptcha(zrpc.MustNewClient(c.CaptchaRpc)),
	}
}

func newGorm(dataSource string) (*gorm.DB, error) {
	// TODO: 利用gorm的DBResolver实现读写分离
	return gorm.Open(mysql.Open(dataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "ttm_",
			SingularTable: true,
		},
	})
}
