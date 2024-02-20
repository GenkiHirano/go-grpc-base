package schema

import "github.com/uptrace/bun"

// TODO: indexを追加する
type SampleUser struct {
	bun.BaseModel `bun:"table:sample_user,alias:u"`
	ID            int64 `bun:",pk,autoincrement"`
	Name          string
}
