package model

import (
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		new(User),
		new(UserAuth),
	)
}
