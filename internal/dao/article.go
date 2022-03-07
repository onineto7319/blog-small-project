package dao

import (
	"database/sql"
	"time"

	"github.com/blog-small-project/internal/dto"
)

type article struct {
	db *sql.DB
}

type ArticleDaoInterface interface {
	Get(uint32) (*dto.Article, error)
	GetAll() ([]dto.Article, error)
	GetAllWithState(int8) ([]dto.Article, error)
	Create(dto.Article) (int64, error)
	Update(dto.Article) error
	Delete(uint32, string) error
}

func NewArticleDao(mysql *sql.DB) ArticleDaoInterface {
	return &article{db: mysql}
}

func (a *article) Get(articleid uint32) (*dto.Article, error) {
	err := a.db.Ping()
	if err != nil {
		return nil, err
	}

	sqlGetString := `SELECT id, title, blog_article.desc, conver_image_url, content, state, created_on, created_by, modified_on, modified_by
					 FROM blog_article
					 WHERE id = ? AND is_del = 0;`

	queryRes := a.db.QueryRow(sqlGetString, articleid)

	res := dto.Article{}
	err = queryRes.Scan(&res.ID, &res.Title, &res.Desc, &res.ConverImageUrl, &res.Content, &res.State,
		&res.CreatedOn, &res.CreatedBy, &res.ModifiedOn, &res.ModifiedBy)

	if err != nil {
		return nil, err
	}
	return &res, nil
}
func (a *article) GetAll() ([]dto.Article, error) {
	err := a.db.Ping()
	if err != nil {
		return nil, err
	}

	sqlGetString := `SELECT id, title, blog_article.desc, conver_image_url, content, state, created_on, created_by, modified_on, modified_by
					 FROM blog_article
					 WHERE is_del = 0;`

	queryRes, err := a.db.Query(sqlGetString)

	if err != nil {
		return nil, err
	}
	res := make([]dto.Article, 0)

	for queryRes.Next() {
		temp := dto.Article{}
		err := queryRes.Scan(&temp.ID, &temp.Title, &temp.Desc, &temp.ConverImageUrl, &temp.Content, &temp.State,
			&temp.CreatedOn, &temp.CreatedBy, &temp.ModifiedOn, &temp.ModifiedBy)
		if err != nil {
			return nil, err
		}
		res = append(res, temp)
	}

	return res, nil
}
func (a *article) GetAllWithState(state int8) ([]dto.Article, error) {
	err := a.db.Ping()
	if err != nil {
		return nil, err
	}

	sqlGetString := `SELECT id, title, desc, conver_image_url, content, state, created_on, created_by, modified_on, modified_by
					 FROM blog_article
					 WHERE is_del = 0 AND state = ?;`

	queryRes, err := a.db.Query(sqlGetString, state)

	if err != nil {
		return nil, err
	}
	res := make([]dto.Article, 0)

	for queryRes.Next() {
		temp := dto.Article{}
		err := queryRes.Scan(&temp.ID, &temp.Title, &temp.Desc, &temp.ConverImageUrl, &temp.Content, &temp.State,
			&temp.CreatedOn, &temp.CreatedBy, &temp.ModifiedOn, &temp.ModifiedBy)
		if err != nil {
			return nil, err
		}
		res = append(res, temp)
	}

	return res, nil
}

func (a *article) Create(article dto.Article) (int64, error) {
	err := a.db.Ping()
	if err != nil {
		return -1, err
	}

	sqlcreateString := `INSERT INTO blog_article (title, blog_article.desc, conver_image_url, content, modified_on, modified_by)
						VALUES (?, ?, ?, ?, ?, ?);`

	res, err := a.db.Exec(sqlcreateString, article.Title, article.Desc, article.ConverImageUrl, article.Content, time.Now().Unix(), article.ModifiedBy)
	if err != nil {
		return -1, err
	}

	resId, _ := res.LastInsertId()

	return resId, nil
}
func (a *article) Update(article dto.Article) error {
	err := a.db.Ping()
	if err != nil {
		return err
	}

	sqlupdatestring := `UPDATE blog_article
						SET title = ?, blog_article.desc = ?, conver_image_url= ?, content = ?, state = ?, modified_on = ?, modified_by = ?
						WHERE id = ?`

	_, err = a.db.Exec(sqlupdatestring, article.Title, article.Desc, article.ConverImageUrl, article.Content, article.State, time.Now().Unix(), article.ModifiedBy, article.ID)

	if err != nil {
		return err
	}

	return nil
}
func (a *article) Delete(articleid uint32, modifiedBy string) error {
	err := a.db.Ping()
	if err != nil {
		return err
	}

	sqldeletestring := `UPDATE blog_article 
						SET is_del = 1, modified_on = ?, modified_by = ?, deleted_on = ?
						WHERE id= ?;`

	timeNow := time.Now().Unix()
	_, err = a.db.Exec(sqldeletestring, timeNow, modifiedBy, timeNow, articleid)

	if err != nil {
		return err
	}

	return nil
}
