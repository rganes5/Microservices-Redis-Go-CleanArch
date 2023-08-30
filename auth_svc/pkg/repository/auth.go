package repository

import (
	"X-TENTIONCREW/auth_svc/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type authRepo struct {
	DB *gorm.DB
}

func NewauthRepo(db *gorm.DB) interfaces.AuthRepo {
	return &authRepo{
		DB: db,
	}
}
