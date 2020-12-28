package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
)

type DBLogger struct{}

// 实现pg框架的一个日志钩子
func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

func New(opts *pg.Options) *pg.DB {
	return pg.Connect(opts)
}
