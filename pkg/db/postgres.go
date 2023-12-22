package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Spargwy/just-to-do-it/pkg/logger"
	"github.com/go-pg/pg/v10"
)

type DBConfig struct {
	DBURL           string
	EnableDBLog     bool
	EnableLongDBLog bool
	MaxRetries      int
}

func NewPostgres(cfg DBConfig) (*pg.DB, error) {
	opts, err := pg.ParseURL(cfg.DBURL)
	if err != nil {
		return nil, fmt.Errorf("parseURL: %v", err)
	}
	opts.MaxRetries = cfg.MaxRetries

	db := pg.Connect(opts)
	if cfg.EnableDBLog {
		db.AddQueryHook(dbLogger{})
	}

	if cfg.EnableLongDBLog {
		db.AddQueryHook(longQueryLogger{})
	}

	err = db.Ping(context.Background())

	return db, err
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d dbLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	sql, err := q.FormattedQuery()
	mes := string(sql)
	logger.Info(mes, err)

	return nil
}

type longQueryLogger struct{}

func (d longQueryLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d longQueryLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	diff := time.Since(q.StartTime)
	if diff.Seconds() > 1 {
		sql, err := q.FormattedQuery()
		log.Println("detect slow query", string(sql), err, diff.Seconds())
	}
	return nil
}
