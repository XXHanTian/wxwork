package type_kf

type AddContactWayResp struct {
	BaseModel
	Url string `json:"url"`
}

type SendMessageResp struct {
	BaseModel
	MsgId string `json:"msgid"`
}

type TransferStateResp struct {
	BaseModel
	MsgCode string `json:"msg_code"`
}
