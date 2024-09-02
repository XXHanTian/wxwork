package type_kf

type BaseModel struct {
	ErrCode int64  `json:"errcode"` // 出错返回码，为0表示成功，非0表示调用失败
	ErrMsg  string `json:"errmsg"`  // 返回码提示语
}
