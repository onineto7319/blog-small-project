package model

type GetArticleRequest struct {
	ID    uint32 `json:"id"`
	State int8   `json:"state"`
}

type GetArticleResponse struct {
	ID             uint32 `json:"id"`
	Title          string `json:"title"`
	Desc           string `json:"desc"`
	ConverImageUrl string `json:"conver_image_url"`
	Content        string `json:"content"`
	State          uint8  `json:"state"`
	CreatedBy      string `json:"created_by"`
	CreatedOn      uint32 `json:"created_on"`
	ModifiedBy     string `json:"modified_by"`
	ModifiedOn     uint32 `json:"modifed_on"`
}

type CreateArticleRequest struct {
	Title          string `json:"title"`
	Desc           string `json:"desc"`
	ConverImageUrl string `json:"conver_image_url"`
	Content        string `json:"content"`
	CreatedBy      string `json:"created_by"`
}

type CreateArticleResponse struct {
	ID             uint32 `json:"id"`
	Title          string `json:"title"`
	Desc           string `json:"desc"`
	ConverImageUrl string `json:"conver_image_url"`
	Content        string `json:"content"`
	State          uint8  `json:"state"`
	CreatedBy      string `json:"created_by"`
	CreatedOn      uint32 `json:"created_on"`
}

type UpdateArticleRequest struct {
	ID             uint32 `json:"id"`
	Title          string `json:"title"`
	Desc           string `json:"desc"`
	ConverImageUrl string `json:"conver_image_url"`
	Content        string `json:"content"`
	State          uint8  `json:"state"`
	ModifiedBy     string `json:"modified_by"`
}

type UpdateArticleResponse struct {
	ID             uint32 `json:"id"`
	Title          string `json:"title"`
	Desc           string `json:"desc"`
	ConverImageUrl string `json:"conver_image_url"`
	Content        string `json:"content"`
	State          uint8  `json:"state"`
	CreatedBy      string `json:"created_by"`
	CreatedOn      uint32 `json:"created_on"`
	ModifiedBy     string `json:"modified_by"`
	ModifiedOn     uint32 `json:"modifed_on"`
}

type DeleteArticleRequest struct {
	ID         uint32 `json:"id"`
	ModifiedBy string `json:"modified_by"`
}

type DeleteArticleResponse struct {
	ID    uint32 `json:"id"`
	IsDel bool   `json:"is_del"`
}
