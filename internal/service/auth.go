package service

import (
	"database/sql"
	"errors"

	"github.com/blog-small-project/internal/dao"

	"github.com/blog-small-project/internal/model"
)

type auth struct {
	db *sql.DB
}

type AuthInterface interface {
	CheckAuth(model.CheckAuthRequest) error
}

func NewAuth(mysql *sql.DB) AuthInterface {
	return &auth{db: mysql}
}

func (a *auth) CheckAuth(requestModel model.CheckAuthRequest) error {
	dto, err := dao.NewAuthDao(a.db).Get(requestModel.AppKey, requestModel.AppSecret)

	if err != nil {
		return err
	}

	if dto.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist.")
}
