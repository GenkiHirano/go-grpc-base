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
		// TODO: カスタムエラーでラップする
		return err
	}

	db, err := cfg.DB.Init()
	if err != nil {
		//　TODO: カスタムエラーでラップする
		return err
	}

	defer db.Close()

	for _, v := range schema.Schemas {
		query := db.NewCreateTable().Model(v).IfNotExists()
		rawQuery, err := query.AppendQuery(db.Formatter(), nil)
		if err != nil {
			fmt.Println("🔥 AppendQuery Error: ", err)
			return err
		}

		fileName := query.GetTableName() + ".sql"
		filePath := filepath.Join(".", fileName)

		if err := os.WriteFile(filePath, rawQuery, 0644); err != nil {
			fmt.Println("🔥 WriteFile Error: ", err)
			return err
		}

		if _, err := db.ExecContext(ctx, string(rawQuery)); err != nil {
			fmt.Println("🔥 ExecContext Error: ", err)
			return err
		}

		fmt.Println("テーブル作成成功: ", query.GetTableName())
	}

	return nil
}
