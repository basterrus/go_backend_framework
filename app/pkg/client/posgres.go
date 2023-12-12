package client

import (
	"context"
	"github.com/basterrus/go_backend_framework/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type PgClient interface {
	Close()
	Acquire(ctx context.Context) (*pgxpool.Conn, error)
	AcquireFunc(ctx context.Context, f func(*pgxpool.Conn) error) error
	AcquireAllIdle(ctx context.Context) []*pgxpool.Conn
	Stat() *pgxpool.Stat
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

func NewPostgresClient(ctx context.Context, cfg *config.Config, connectionString string) (pool *pgxpool.Pool, err error) {
	pgxCfg, parseConfigErr := pgxpool.ParseConfig(connectionString)
	if parseConfigErr != nil {
		log.Printf("Unable to parse config: %v\n", parseConfigErr)
		return nil, parseConfigErr
	}

	if cfg.IsBinary {
		pgxCfg.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe
	}

	pool, parseConfigErr = pgxpool.NewWithConfig(ctx, pgxCfg)
	if parseConfigErr != nil {
		log.Printf("Failed to parse PostgreSQL configuration due to error: %v\n", parseConfigErr)
		return nil, parseConfigErr
	}

	err = DoWithAttempts(func() error {
		pingErr := pool.Ping(ctx)
		if pingErr != nil {
			log.Printf("Failed to connect to postgres due to error %v... Going to do the next attempt\n", pingErr)
			return pingErr
		}

		return nil
	}, cfg.MaxAttempts, time.Duration(cfg.MaxDelay))
	if err != nil {
		log.Fatal("All attempts are exceeded. Unable to connect to PostgreSQL")
	}

	return pool, nil
}

func DoWithAttempts(fn func() error, maxAttempts int, delay time.Duration) (err error) {
	for maxAttempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			maxAttempts--

			continue
		}
		return nil
	}
	return err
}
