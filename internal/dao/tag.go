package dao

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/blog-small-project/internal/dto"
)

type tag struct {
	db *sql.DB
}

type TagDaoInterface interface {
	Get(uint32) (*dto.Tag, error)
	GetAll() ([]dto.Tag, error)
	GetAllWithState(int8) ([]dto.Tag, error)
	Create(dto.Tag) (int64, error)
	Update(dto.Tag) error
	Delete(uint32, string) error
}

func NewTagDao(mysql *sql.DB) TagDaoInterface {
	return &tag{db: mysql}
}

func (t *tag) Get(tagid uint32) (*dto.Tag, error) {
	err := t.db.Ping()
	if err != nil {
		return nil, err
	}
	sqlquerystring := `SELECT id, name, state, created_on, created_by, modified_on, modified_by
					   FROM blog_tag
					   WHERE id= ? AND is_del = 0;`

	result := dto.Tag{}
	err = t.db.QueryRow(sqlquerystring, tagid).Scan(&result.ID, &result.Name, &result.State,
		&result.CreatedOn, &result.CreatedBy, &result.ModifiedOn, &result.ModifiedBy)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (t *tag) GetAll() ([]dto.Tag, error) {
	err := t.db.Ping()

	if err != nil {
		return nil, err
	}
	sqlqueryString := `SELECT id, name, state, created_on, created_by, modified_on, modified_by
					   FROM blog_tag
					   WHERE is_del = 0;`

	queryRes, err := t.db.Query(sqlqueryString)

	if err != nil {
		return nil, err
	}

	result := make([]dto.Tag, 0)

	for queryRes.Next() {
		temp := dto.Tag{}
		err := queryRes.Scan(&temp.ID, &temp.Name, &temp.State,
			&temp.CreatedOn, &temp.CreatedBy, &temp.ModifiedOn, &temp.ModifiedBy)
		if err != nil {
			return nil, err
		}
		result = append(result, temp)
	}

	return result, nil
}

func (t *tag) GetAllWithState(state int8) ([]dto.Tag, error) {
	err := t.db.Ping()

	if err != nil {
		return nil, err
	}

	sqlqueryString := `SELECT id, name, state, created_on, created_by, modified_on, modified_by
						  FROM blog_tag
						  WHERE state = ?`

	queryRes, err := t.db.Query(sqlqueryString, state)

	if err != nil {
		return nil, err
	}

	result := make([]dto.Tag, 0)

	for queryRes.Next() {
		temp := dto.Tag{}
		err := queryRes.Scan(&temp.ID, &temp.Name, &temp.State,
			&temp.CreatedOn, &temp.CreatedBy, &temp.ModifiedOn, &temp.ModifiedBy)
		if err != nil {
			return nil, err
		}
		result = append(result, temp)
	}

	return result, nil
}

func (t *tag) Create(tag dto.Tag) (int64, error) {
	err := t.db.Ping()
	if err != nil {
		return -1, err
	}

	sqlcreateString := `INSERT INTO blog_tag (name, created_on, created_by)
						VALUES (?, ?, ?);`

	res, err := t.db.Exec(sqlcreateString, tag.Name, time.Now().Unix(), tag.CreatedBy)

	if err != nil {
		return -1, err
	}

	resId, _ := res.LastInsertId()

	return resId, nil
}
func (t *tag) Update(tag dto.Tag) error {
	err := t.db.Ping()
	if err != nil {
		return err
	}
	fmt.Print(tag)
	sqlupdatestring := `UPDATE blog_tag
						SET name = ?, state = ?, modified_on = ?, modified_by = ?
						WHERE id = ?`

	_, err = t.db.Exec(sqlupdatestring, tag.Name, tag.State, time.Now().Unix(), tag.ModifiedBy, tag.ID)

	if err != nil {
		return err
	}
	return nil
}
func (t *tag) Delete(tagid uint32, modifiedBy string) error {
	err := t.db.Ping()
	if err != nil {
		return err
	}

	sqldeletestring := `UPDATE blog_tag 
						SET is_del = 1, modified_on = ?, modified_by = ?, deleted_on = ?
						WHERE id= ?;`

	timeNow := time.Now().Unix()
	_, err = t.db.Exec(sqldeletestring, timeNow, modifiedBy, timeNow, tagid)

	if err != nil {
		return err
	}

	return nil
}
