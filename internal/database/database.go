package database

import (
	"cloud.google.com/go/cloudsqlconn"
	"context"
	"database/sql"
	"fmt"
	"github.com/Vaansh/gore/internal/config"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"net"
)

func ConnectDb() (*sql.DB, error) {
	cfg, err := config.ReadDbConfig()

	var (
		dbUser                 = cfg.Username
		dbPwd                  = cfg.Password
		dbName                 = cfg.Database
		instanceConnectionName = cfg.InstanceId
	)

	dsn := fmt.Sprintf("user=%s password=%s database=%s", dbUser, dbPwd, dbName)
	dbcfg, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	var opts []cloudsqlconn.Option
	d, err := cloudsqlconn.NewDialer(context.Background(), opts...)
	if err != nil {
		return nil, err
	}

	dbcfg.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
		return d.Dial(ctx, instanceConnectionName)
	}

	dbURI := stdlib.RegisterConnConfig(dbcfg)
	dbPool, err := sql.Open("pgx", dbURI)
	dbPool.SetMaxOpenConns(10)

	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	return dbPool, nil
}
