package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gameloungee/server/config"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", config.New().DBConnectionString)

	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxIdleTime(0)
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db
}

func WithPrepare(sql string, args ...any) error {
	conn := Connect()
	defer conn.Close()

	stmt, err := conn.Prepare(sql)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)

	if err != nil {
		return err
	}

	return nil
}

func Get(sqll string, arg any) (*sql.Rows, error) {
	conn := Connect()
	defer conn.Close()

	query := fmt.Sprintf(sqll, arg)
	rows, err := conn.Query(query, arg)

	if err != nil {
		return &sql.Rows{}, err
	}

	return rows, nil
}
