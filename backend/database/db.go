package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/GenkiHirano/go-grpc-base/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/volatiletech/sqlboiler/boil"
)

func Init(cfg config.DBConfig) (*bun.DB, error) {
	sqlDB, err := sql.Open(
		cfg.Driver,
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
		),
	)
	if err != nil {
		// TODO: カスタムエラーでラップする
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		// TODO: カスタムエラーでラップする
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(300 * time.Second)
	boil.SetDB(sqlDB)

	return bun.NewDB(sqlDB, mysqldialect.New()), nil
}
