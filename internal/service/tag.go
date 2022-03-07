package service

import (
	"database/sql"
	"encoding/json"

	"github.com/blog-small-project/internal/dto"

	"github.com/blog-small-project/internal/dao"
	"github.com/blog-small-project/internal/model"
)

type tag struct {
	db *sql.DB
}

type TagInterface interface {
	Get(model.GetTagRequest) (model.GetTagResponse, error)
	GetAll(model.GetTagRequest) ([]model.GetTagResponse, error)
	Create(model.CreateTagRequest) (model.CreateTagResponse, error)
	Update(model.UpdateTagRequest) (model.UpdateTagResponse, error)
	Delete(model.DeleteTagRequest) error
}

func NewTag(mysqlDB *sql.DB) TagInterface {
	return &tag{db: mysqlDB}
}

func (t *tag) Get(requestmodel model.GetTagRequest) (model.GetTagResponse, error) {
	dto, err := dao.NewTagDao(t.db).Get(requestmodel.ID)
	req := model.GetTagResponse{}
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

func (t *tag) GetAll(requestmodel model.GetTagRequest) ([]model.GetTagResponse, error) {
	var (
		dto []dto.Tag
		err error
	)

	if requestmodel.State == -1 {
		dto, err = dao.NewTagDao(t.db).GetAll()
	} else {
		dto, err = dao.NewTagDao(t.db).GetAllWithState(requestmodel.State)
	}

	if err != nil {
		return nil, err
	}

	res := make([]model.GetTagResponse, 0)

	for i := 0; i < len(dto); i++ {
		dtoString, err := json.Marshal(dto[i])
		if err != nil {
			return nil, err
		}
		temp := model.GetTagResponse{}
		err = json.Unmarshal(dtoString, &temp)
		if err != nil {
			return nil, err
		}
		res = append(res, temp)
	}

	return res, nil
}

func (t *tag) Create(requestmodel model.CreateTagRequest) (model.CreateTagResponse, error) {
	tagdao := dao.NewTagDao(t.db)
	createID, err := tagdao.Create(dto.Tag{Name: requestmodel.Name})
	req := model.CreateTagResponse{}
	if err != nil {
		return req, err
	}
	dto, err := tagdao.Get(uint32(createID))

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

func (t *tag) Update(requestmodel model.UpdateTagRequest) (model.UpdateTagResponse, error) {
	tagdao := dao.NewTagDao(t.db)
	err := tagdao.Update(dto.Tag{Name: requestmodel.Name, State: int8(requestmodel.State), Common: dto.Common{ID: requestmodel.ID, ModifiedBy: requestmodel.ModifiedBy}})
	req := model.UpdateTagResponse{}
	if err != nil {
		return req, err
	}
	getDto, err := tagdao.Get(requestmodel.ID)

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

func (t *tag) Delete(requestmodel model.DeleteTagRequest) error {
	return dao.NewTagDao(t.db).Delete(requestmodel.ID, requestmodel.ModifiedBy)
}
