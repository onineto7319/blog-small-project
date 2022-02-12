package dao

import (
	"database/sql"
	"fmt"

	"github.com/blog-small-project/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLDBEngine(databaseMysqlSetting *setting.DatabaseMysqlSettings) (*sql.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseMysqlSetting.Username,
		databaseMysqlSetting.Password,
		databaseMysqlSetting.Host,
		databaseMysqlSetting.DBName,
		databaseMysqlSetting.Charset,
		databaseMysqlSetting.ParseTime)

	db, err := sql.Open(databaseMysqlSetting.DBType, s)

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(databaseMysqlSetting.MaxIdleConns)
	db.SetMaxOpenConns(databaseMysqlSetting.MaxOpenConns)
	return db, err
}
