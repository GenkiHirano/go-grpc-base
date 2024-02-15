package migrate

import (
	"context"

	"github.com/GenkiHirano/go-grpc-base/database/schema"
	"github.com/uptrace/bun"
)

func ResetTables(ctx context.Context, db *bun.DB) error {
	for _, v := range schema.Schemas {
		if err := db.ResetModel(ctx, v); err != nil {
			// カスタムエラーでラップするか、適切なエラーメッセージを表示
			return err
		}
	}
	return nil
}
