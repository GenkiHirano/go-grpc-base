package migrate

import (
	"context"

	"github.com/GenkiHirano/go-grpc-base/database/schema"
	"github.com/uptrace/bun"
)

func CreateTables(ctx context.Context, db *bun.DB) error {
	for _, v := range schema.Schemas {
		if _, err := db.NewCreateTable().Model(v).IfNotExists().Exec(ctx); err != nil {
			// TODO: カスタムエラーでラップする
			return err
		}
	}
	return nil
}
