package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"grpcDbSync/api/grpcapp"
	"grpcDbSync/store/repositories/Postgres/DbSyncPg"
	"os"
)

func main() {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	err = pool.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	DbSyncRepository := DbSyncPg.New(pool)

	application := grpcapp.New(8980, DbSyncRepository)
	err = application.Start()
	if err != nil {
		return
	}
}
