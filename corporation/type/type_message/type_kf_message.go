package type_message

type SyncMsgSchema struct {
	ErrCode    int32        `json:"errcode"`     // 返回码
	ErrMsg     string       `json:"errmsg"`      // 错误码描述
	NextCursor string       `json:"next_cursor"` // 下次调用带上该值，则从当前的位置继续往后拉，以实现增量拉取。强烈建议对改该字段入库保存，每次请求读取带上，请求结束后更新。避免因意外丢，导致必须从头开始拉取，引起消息延迟。
	HasMore    uint32       `json:"has_more"`    // 是否还有更多数据。0-否；1-是。不能通过判断msg_list是否空来停止拉取，可能会出现has_more为1，而msg_list为空的情况
	MsgList    []*KfMessage `json:"msg_list"`    // 消息列表
}

type SyncMsgSchemaRaw struct {
	ErrCode    int32                    `json:"errcode"`     // 返回码
	ErrMsg     string                   `json:"errmsg"`      // 错误码描述
	NextCursor string                   `json:"next_cursor"` // 下次调用带上该值，则从当前的位置继续往后拉，以实现增量拉取。强烈建议对改该字段入库保存，每次请求读取带上，请求结束后更新。避免因意外丢，导致必须从头开始拉取，引起消息延迟。
	HasMore    uint32                   `json:"has_more"`    // 是否还有更多数据。0-否；1-是。不能通过判断msg_list是否空来停止拉取，可能会出现has_more为1，而msg_list为空的情况
	MsgList    []map[string]interface{} `json:"msg_list"`    // 消息列表
}

// Message 同步的消息内容
type KfMessage struct {
	MsgID              string `json:"msgid"`           // 消息ID
	OpenKFID           string `json:"open_kfid"`       // 客服帐号ID（msgtype为event，该字段不返回）
	ExternalUserID     string `json:"external_userid"` // 客户UserID（msgtype为event，该字段不返回）
	ReceptionistUserID string `json:"servicer_userid"` // 接待客服userID
	SendTime           uint64 `json:"send_time"`       // 消息发送时间
	Origin             uint32 `json:"origin"`          // 消息来源。3-微信客户发送的消息 4-系统推送的事件消息 5-接待人员在企业微信客户端发送的消息
	MsgType            string `json:"msgtype"`         // 消息类型
	EventType          string `json:"event_type"`      // 事件类型
	OriginData         []byte `json:"origin_data"`     // 原始数据内容
}
