package db

import (
	"log"

	"github.com/hrvadl/studdy-buddy/auth/pkg/config"
	"github.com/hrvadl/studdy-buddy/auth/pkg/models"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func Init(config *config.Config) *gorm.DB {
	mysqlConn := mysql.Open(config.DNS)
	db, err := gorm.Open(mysqlConn, &gorm.Config{})

	if err != nil {
		log.Fatalf("cannot open connection tp DB: %v", err)
	}

	db.AutoMigrate(&models.User{})

	return db
}
