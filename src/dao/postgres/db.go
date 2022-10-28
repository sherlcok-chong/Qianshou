package postgres

import (
	db "Qianshou/src/dao/postgres/sqlc"
	"Qianshou/src/dao/postgres/tx"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB interface {
	tx.TXer
	db.Querier
}

func Init(dataSourceName string) DB {
	pool, err := pgxpool.Connect(context.Background(), dataSourceName)
	if err != nil {
		panic(err)
	}
	return &tx.SqlStore{Queries: db.New(pool), DB: pool}
}
