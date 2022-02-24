package service

import (
	"database/sql"
	"encoding/json"

	"github.com/blog-small-project/internal/dao"
	"github.com/blog-small-project/internal/dto"
	"github.com/blog-small-project/internal/model"
)

type article struct {
	db *sql.DB
}

type ArticleInterface interface {
	Get(model.GetArticleRequest) (model.GetArticleResponse, error)
	GetAll(model.GetArticleRequest) ([]model.GetArticleResponse, error)
	Create(model.CreateArticleRequest) (model.CreateArticleResponse, error)
	Update(model.UpdateArticleRequest) (model.UpdateArticleResponse, error)
	Delete(model.DeleteArticleRequest) error
}

func NewArticle(mysql *sql.DB) ArticleInterface {
	return &article{db: mysql}
}

func (a *article) Get(requestmodel model.GetArticleRequest) (model.GetArticleResponse, error) {
	dto, err := dao.NewArticleDao(a.db).Get(requestmodel.ID)
	req := model.GetArticleResponse{}
	if err != nil {
		return req, err
	}

	dtoString, err := json.Marshal(dto)

	if err != nil {
		return req, err
	}

	err = json.Unmarshal(dtoString, &req)

	if err != nil {
		return req, err
	}

	return req, nil
}
func (a *article) GetAll(requestmodel model.GetArticleRequest) ([]model.GetArticleResponse, error) {
	var (
		dto []dto.Article
		err error
	)

	if requestmodel.State == -1 {
		dto, err = dao.NewArticleDao(a.db).GetAll()
	} else {
		dto, err = dao.NewArticleDao(a.db).GetAllWithState(requestmodel.State)
	}

	if err != nil {
		return nil, err
	}

	res := make([]model.GetArticleResponse, 0)

	for i := 0; i < len(dto); i++ {
		dtoString, err := json.Marshal(dto[i])
		if err != nil {
			return nil, err
		}
		temp := model.GetArticleResponse{}
		err = json.Unmarshal(dtoString, &temp)
		if err != nil {
			return nil, err
		}
		res = append(res, temp)
	}

	return res, nil
}
func (a *article) Create(requestmodel model.CreateArticleRequest) (model.CreateArticleResponse, error) {
	articledao := dao.NewArticleDao(a.db)
	createID, err := articledao.Create(dto.Article{Title: requestmodel.Title, Desc: requestmodel.Desc, ConverImageUrl: requestmodel.ConverImageUrl, Content: requestmodel.Content})
	req := model.CreateArticleResponse{}
	if err != nil {
		return req, err
	}
	dto, err := articledao.Get(uint32(createID))

	if err != nil {
		return req, err
	}
	dtoString, err := json.Marshal(dto)
	if err != nil {
		return req, err
	}
	err = json.Unmarshal(dtoString, &req)
	if err != nil {
		return req, err
	}

	return req, nil
}
func (a *article) Update(requestmodel model.UpdateArticleRequest) (model.UpdateArticleResponse, error) {
	articledao := dao.NewArticleDao(a.db)
	err := articledao.Update(dto.Article{Title: requestmodel.Title, Desc: requestmodel.Desc, ConverImageUrl: requestmodel.ConverImageUrl, Content: requestmodel.Content, State: requestmodel.State, Common: dto.Common{ID: requestmodel.ID, ModifiedBy: requestmodel.ModifiedBy}})
	req := model.UpdateArticleResponse{}
	if err != nil {
		return req, err
	}

	getDto, err := articledao.Get(requestmodel.ID)

	if err != nil {
		return req, err
	}

	dtoString, err := json.Marshal(getDto)

	if err != nil {
		return req, err
	}

	err = json.Unmarshal(dtoString, &req)
	if err != nil {
		return req, err
	}

	return req, nil
}
func (a *article) Delete(requestmodel model.DeleteArticleRequest) error {
	return dao.NewArticleDao(a.db).Delete(requestmodel.ID, requestmodel.ModifiedBy)
}
