package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"os"
)

func EstablishPSQLConnection() (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), fmt.Sprintf("%s",
		os.Getenv("POSTGRESQL_URL")))
	if err != nil {
		logrus.Fatal(err)
	}

	return db, nil
}

func CloseConn(db *pgx.Conn) error {
	return db.Close(context.Background())
}
