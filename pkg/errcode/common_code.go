package errcode

var (
	Success       = NewError(0, "成功")
	ServerError   = NewError(10000000, "服務內部錯誤")
	InvalidParams = NewError(10000001, "導入參數錯誤")
	NotFound      = NewError(10000002, "找不到")
)
