package svc

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/yguilai/timetable-micro/model"
	"github.com/yguilai/timetable-micro/services/user/rpc/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	c  config.Config
	Db *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "ttm_",
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		new(model.User),
		new(model.UserAuth),
	)
	if err != nil {
		logx.Error(err)
	}

	return &ServiceContext{
		c:  c,
		Db: db,
	}
}
