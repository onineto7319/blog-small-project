package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服務內部錯誤")
	InvalidParams             = NewError(10000001, "導入參數錯誤")
	NotFound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "驗證失敗，找不到對應的AppKey和AppSecret")
	UnauthorizedNotToken      = NewError(10000004, "驗證失敗，未輸入Token")
	UnauthorizedTokenError    = NewError(10000005, "驗證失敗，Token錯誤")
	UnauthorizedTokenTimeout  = NewError(10000006, "驗證失敗 Token逾時")
	UnauthorizedTokenGenerate = NewError(10000007, "驗證失敗，Token產生失敗")
)
