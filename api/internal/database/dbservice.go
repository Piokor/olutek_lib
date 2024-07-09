package database

import (
	"github.com/jmoiron/sqlx"
)

type DbService struct {
	client *sqlx.DB
}

func NewDbService() (*DbService, error) {
	dbConnection, err := getClient()
	if err != nil {
		return nil, err
	}
	service := DbService{dbConnection}
	return &service, nil
}
