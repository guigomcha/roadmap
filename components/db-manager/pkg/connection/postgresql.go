package connection

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToPostgreSQL() (*gorm.DB, error) {
	// https://github.com/go-gorm/postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
