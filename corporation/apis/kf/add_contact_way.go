package kf

import (
	"bytes"
	"encoding/json"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/type/type_kf"
)

const (
	apiKfAddContactWay = "/cgi-bin/kf/add_contact_way"
)

func KfAddContactWay(ctx *corporation.App, openKfId string, scene string) (*type_kf.AddContactWayResp, error) {
	m := map[string]interface{}{
		"open_kfid": openKfId,
		"scene":     scene,
	}
	payload, _ := json.Marshal(m)
	data, err := ctx.Client.HTTPPost(apiKfAddContactWay, bytes.NewReader(payload), "application/json;charset=utf-8")
	if err != nil {
		return nil, err
	}

	rsp := type_kf.AddContactWayResp{}
	err = json.Unmarshal(data, &rsp)
	return &rsp, err
}
