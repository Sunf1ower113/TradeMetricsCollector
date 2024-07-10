package postgreSQL

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"tradeMetricsCollector/internal/config"
)

func ConnectToDB(cfg *config.Config) (db *sql.DB, err error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
