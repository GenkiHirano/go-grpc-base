package schema

import "github.com/uptrace/bun"

// TODO: indexを追加する
type SampleCompany struct {
	bun.BaseModel `bun:"table:sample_company,alias:u"`
	ID            int64 `bun:",pk,autoincrement"`
	Name          string
}
