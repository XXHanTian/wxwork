package kf

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiKfSendMessageOnEvent = "/cgi-bin/kf/send_msg_on_event"
)

func SendMessageOnEvent(ctx *corporation.App, payload []byte) ([]byte, error) {
	data, err := ctx.Client.HTTPPost(apiKfSendMessageOnEvent, bytes.NewReader(payload), "application/json;charset=utf-8")
	return data, err
}
