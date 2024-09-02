package kf

import (
	"bytes"
	"encoding/json"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/type/type_kf"
)

const (
	apiSendMessage = "/cgi-bin/kf/send_msg"
)

func SendMessageText(ctx *corporation.App, toUser, openKfid string, content string) (*type_kf.SendMessageResp, error) {
	params := map[string]interface{}{
		"touser":    toUser,
		"open_kfid": openKfid,
		"msgtype":   "text",
		"text": map[string]interface{}{
			"content": content,
		},
	}
	payload, _ := json.Marshal(params)
	data, err := ctx.Client.HTTPPost(apiSendMessage, bytes.NewReader(payload), "application/json;charset=utf-8")
	if err != nil {
		return nil, err
	}

	rsp := type_kf.SendMessageResp{}
	err = json.Unmarshal(data, &rsp)
	return &rsp, err
}
