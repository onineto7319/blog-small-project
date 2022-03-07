package dao

import (
	"database/sql"

	"github.com/blog-small-project/internal/dto"
)

type auth struct {
	db *sql.DB
}

type AuthDaoInterface interface {
	Get(app_key, app_secret string) (*dto.Auth, error)
}

func NewAuthDao(mysqlDB *sql.DB) AuthDaoInterface {
	return &auth{db: mysqlDB}
}

func (a *auth) Get(app_key, app_secret string) (*dto.Auth, error) {
	err := a.db.Ping()
	if err != nil {
		return nil, err
	}
	sqlGetString := `SELECT id, app_key, app_secret
					 FROM blog_auth
					 WHERE app_key = ? AND app_secret = ?;`

	queryRes := a.db.QueryRow(sqlGetString, app_key, app_secret)
	res := dto.Auth{}

	err = queryRes.Scan(&res.ID, &res.AppKey, &res.AppSecret)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
