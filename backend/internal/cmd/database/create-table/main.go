package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/GenkiHirano/go-grpc-base/internal/config"
	"github.com/GenkiHirano/go-grpc-base/internal/database/schema"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := createTable(ctx); err != nil {
		// TODO: エラーをカスタムログで出力する
		fmt.Println("create table error: ", err)
	} else {
		// TODO: 成功メッセージをカスタムログで出力する
		fmt.Println("create table success")
	}
}

func createTable(ctx context.Context) error {
	cfg, err := config.Init(ctx)
	if err != nil {
		return err
	}

	db, err := cfg.DB.Init()
	if err != nil {
		return err
	}

	defer db.Close()

	for _, v := range schema.Schemas {
		query := db.NewCreateTable().Model(v).IfNotExists()
		rawQuery, err := query.AppendQuery(db.Formatter(), nil)
		if err != nil {
			//　TODO: カスタムエラーでラップする
			return err
		}

		fileName := query.GetTableName() + ".sql"
		filePath := filepath.Join("./internal/database/migrate/create-table/", fileName)

		if err := os.WriteFile(filePath, rawQuery, 0644); err != nil {
			//　TODO: カスタムエラーでラップする
			return err
		}

		if _, err := db.ExecContext(ctx, string(rawQuery)); err != nil {
			//　TODO: カスタムエラーでラップする
			return err
		}

		// TODO: ログで出力する
		fmt.Println("テーブル作成成功: ", query.GetTableName())
	}

	return nil
}
