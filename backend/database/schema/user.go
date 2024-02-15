package schema

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:user,alias:u"`
	ID            int64 `bun:",pk,autoincrement"`
	Name          string
}
