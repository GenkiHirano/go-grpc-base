package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/GenkiHirano/go-grpc-base/internal/config"
	"github.com/GenkiHirano/go-grpc-base/internal/database/schema"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := dropTable(ctx); err != nil {
		// TODO: エラーをカスタムログで出力する
		fmt.Println("drop table error: ", err)
	} else {
		// TODO: 成功メッセージをカスタムログで出力する
		fmt.Println("drop table success")
	}
}

func dropTable(ctx context.Context) error {
	cfg, err := config.Init(ctx)
	if err != nil {
		// TODO: カスタムエラーでラップする
		return err
	}

	db, err := cfg.DB.Init()
	if err != nil {
		// TODO: カスタムエラーでラップする
		return err
	}

	defer db.Close()

	for _, v := range schema.Schemas {
		query := db.NewDropTable().Model(v).IfExists()
		rawQuery, err := query.AppendQuery(db.Formatter(), nil)
		if err != nil {
			// TODO: カスタムエラーでラップする
			return err
		}

		if _, err := db.ExecContext(ctx, string(rawQuery)); err != nil {
			// TODO: カスタムエラーでラップする
			return err
		}

		// TODO: ログで出力する
		fmt.Println("テーブル削除成功: ", query.GetTableName())
	}

	return nil
}
