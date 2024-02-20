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

	if err := resetTable(ctx); err != nil {
		// TODO: エラーをカスタムログで出力する
		fmt.Println("reset table error: ", err)
	} else {
		// TODO: 成功メッセージをカスタムログで出力する
		fmt.Println("reset table success")
	}
}

func resetTable(ctx context.Context) error {
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
		if err := db.ResetModel(ctx, v); err != nil {
			// TODO: カスタムエラーでラップする
			return err
		}
	}

	return nil
}
