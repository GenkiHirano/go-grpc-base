package migrate

import (
	"context"

	"github.com/GenkiHirano/go-grpc-base/database/schema"
	"github.com/uptrace/bun"
)

func DropTables(ctx context.Context, db *bun.DB) error {
	for _, v := range schema.Schemas {
		if _, err := db.NewDropTable().Model(v).IfExists().Exec(ctx); err != nil {
			// カスタムエラーでラップするか、適切なエラーメッセージを表示
			return err
		}
	}
	return nil
}
