package kf

import (
	"bytes"
	"github.com/fastwego/wxwork/kf"
)

const (
	apiKfSendMessageOnEvent = "/cgi-bin/kf/send_msg"
)

func SendMessageOnEvent(ctx *kf.KfApp, payload []byte) ([]byte, error) {
	data, err := ctx.Client.HTTPPost(apiKfSendMessageOnEvent, bytes.NewReader(payload), "application/json;charset=utf-8")
	return data, err
}
