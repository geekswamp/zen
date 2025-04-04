package di

import (
	"github.com/geekswamp/zen/internal/storage/postgres"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var PostgresSet = wire.NewSet(
	InitPostgres,
	InitGorm,
)

func InitPostgres() postgres.Postgres {
	return postgres.NewDefault()
}

func InitGorm(client postgres.Postgres) *gorm.DB {
	db, err := client.Connect()
	if err != nil {
		panic("failed to connect to DB " + err.Error())
	}
	return db
}
