package type_kf

// CustomerSchema 微信客户基本资料
type CustomerSchema struct {
	ExternalUserID      string `json:"external_userid"` // 微信客户的external_userid
	NickName            string `json:"nickname"`        // 微信昵称
	Avatar              string `json:"avatar"`          // 微信头像。第三方不可获取
	Gender              int    `json:"gender"`          // 性别
	UnionID             string `json:"unionid"`         // unionid，需要绑定微信开发者帐号才能获取到，查看绑定方法: https://open.work.weixin.qq.com/kf/doc/92512/93143/94769#%E5%A6%82%E4%BD%95%E8%8E%B7%E5%8F%96%E5%BE%AE%E4%BF%A1%E5%AE%A2%E6%88%B7%E7%9A%84unionid
	EnterSessionContext struct {
		Scene          string `json:"scene"`
		SceneParam     string `json:"scene_param"`
		WechatChannels struct {
			Nickname string `json:"nickname"`
			Scene    int64  `json:"scene"`
		} `json:"wechat_channels"`
	} `json:"enter_session_context"`
}

// CustomerBatchGetSchema 获取客户基本信息响应内容
type CustomerBatchGetSchema struct {
	BaseModel
	CustomerList          []*CustomerSchema `json:"customer_list"`           // 微信客户信息列表
	InvalidExternalUserID []string          `json:"invalid_external_userid"` // 无效的微信客户ID
}
