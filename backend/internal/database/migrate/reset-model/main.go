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

	cfg, err := config.Init(ctx)
	if err != nil {
		fmt.Println("config error: ", err)
		// return err
	}

	db, err := cfg.DB.Init()
	if err != nil {
		fmt.Println("db error: ", err)
		// return err
	}

	defer db.Close()

	for _, v := range schema.Schemas {
		if err := db.ResetModel(ctx, v); err != nil {
			// カスタムエラーでラップするか、適切なエラーメッセージを表示
			// return err
			fmt.Println("db error: ", err)
		}
	}

	// return nil
}
