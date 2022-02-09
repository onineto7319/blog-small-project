package model

type Common struct {
	ID         uint32 `json:"id"`
	CreatedOn  uint32 `json:"created_on"`
	CreatedBy  string `json:"create_by"`
	ModifiedOn uint32 `json:"modified_on"`
	ModifiedBy string `json:"modified_by"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}
