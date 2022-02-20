package model

type GetTagRequest struct {
	ID    uint32 `json:"id"`
	State int8   `json:"state"`
}

type GetTagResponse struct {
	ID         uint32 `json:"id"`
	Name       string `json:"name"`
	State      uint8  `json:"state"`
	CreatedBy  string `json:"created_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedBy string `json:"modified_by"`
	ModifiedOn uint32 `json:"modifed_on"`
}

type CreateTagRequest struct {
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
}

type CreateTagResponse struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	State     uint8  `json:"state"`
	CreatedBy string `json:"created_by"`
	CreatedOn uint32 `json:"created_on"`
}

type UpdateTagRequest struct {
	ID         uint32 `json:"id"`
	Name       string `json:"name"`
	State      uint8  `json:"state"`
	ModifiedBy string `json:"modified_by"`
}

type UpdateTagResponse struct {
	ID         uint32 `json:"id"`
	Name       string `json:"name"`
	State      uint8  `json:"state"`
	CreatedBy  string `json:"created_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedBy string `json:"modified_by"`
	ModifiedOn uint32 `json:"modifed_on"`
}

type DeleteTagRequest struct {
	ID uint32 `json:"id"`
}

type DeleteTagResponse struct {
	ID    int64 `json:"id"`
	IsDel bool  `json:"is_del"`
}
