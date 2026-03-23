package repositories

import (
	"database/sql"
	"fmt"

	"babelbridge/internal/config"
)

type Tables string

const (
	CompanyTable Tables = "company"
	TopicTable   Tables = "topic"
)

type PostgresRepository struct {
	DB *sql.DB
}

func NewPG(cfg *config.Config) *PostgresRepository {
	dbcfg := config.DatabaseConfig{
		Driver:   cfg.DBCfg.Driver,
		Host:     cfg.DBCfg.Host,
		Port:     cfg.DBCfg.Port,
		Username: cfg.DBCfg.Username,
		Name:     cfg.DBCfg.Name,
		Password: cfg.DBCfg.Password,
	}
	var connStr string
	if dbcfg.Driver == "postgres" {
		connStr = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbcfg.Host, dbcfg.Port, dbcfg.Username, dbcfg.Password, dbcfg.Name,
		)
	}
	fmt.Println(connStr)
	db, err := sql.Open(dbcfg.Driver, connStr)
	if err != nil {
		panic(err)
	}

	// if err := db.Ping(); err != nil {
	// 	panic(fmt.Sprintf("failed to connect to db: %v", err))
	// }
	return &PostgresRepository{db}
}
