package database

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func CreateMysqlTables(mysqlDriver *sql.DB) error {

	driver, _ := mysql.WithInstance(mysqlDriver, &mysql.Config{})
	_, err := migrate.NewWithDatabaseInstance("file:./migrations", "mysql", driver)

	if err != nil {
		return err
	}

	return nil
}
