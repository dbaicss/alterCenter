package gorm

import (
	"github.com/jinzhu/gorm"
	"alterCenter/config"
	"alterCenter/models"
)

func AutoMigrate(db *gorm.DB) error {
	if config.DatabaseConfig.Dbtype == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	}
	db.SingularTable(true)
	return db.AutoMigrate(
		new(models.Rule),
	).Error
}
