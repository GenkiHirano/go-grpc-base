package config

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/volatiletech/sqlboiler/boil"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Driver   string `envconfig:"DB_DRIVER" required:"true" default:"mysql"`
	Username string `envconfig:"DB_USERNAME" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true" redact:"true"`
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     string `envconfig:"DB_PORT" required:"true" default:"3306"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}

func (d *DB) Init() (*bun.DB, error) {
	sqlDB, err := sql.Open(
		d.Driver,
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			d.Username, d.Password, d.Host, d.Port, d.Name,
		),
	)
	if err != nil {
		fmt.Println("db-error-1: ", err)
		// TODO: カスタムエラーでラップする
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		// TODO: カスタムエラーでラップする
		fmt.Println("db-error-2: ", err)
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(300 * time.Second)
	boil.SetDB(sqlDB)

	return bun.NewDB(sqlDB, mysqldialect.New()), nil
}
