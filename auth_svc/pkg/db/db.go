package db

import (
	"X-TENTIONCREW/auth_svc/pkg/config"
	"X-TENTIONCREW/auth_svc/pkg/domain"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initdb(c *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(c.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&domain.User{})
	return db, err
}
