package database

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type DbService struct {
	Client *sqlx.DB
}

func NewDbService() (*DbService, error) {
	dbConnection, err := getClient()
	if err != nil {
		return nil, err
	}
	service := DbService{dbConnection}
	return &service, nil
}

func (*DbService) insertString(table string, options string, fields ...string) string {
	fieldsStr := strings.Join(fields, ", ")

	placeholders := make([]string, len(fields))
	for i, field := range fields {
		placeholders[i] = fmt.Sprintf(":%s", field)
	}
	placeholdersStr := strings.Join(placeholders, ", ")

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) %s",
		table,
		fieldsStr,
		placeholdersStr,
		options,
	)
}

func (s *DbService) InsertWithOptions(table string, object interface{}, options string, fields ...string) error {
	insertStr := s.insertString(table, options, fields...)
	_, err := s.Client.NamedExec(insertStr, object)
	return err
}

func (s *DbService) Insert(table string, object interface{}, fields ...string) error {
	return s.InsertWithOptions(table, object, "", fields...)
}
