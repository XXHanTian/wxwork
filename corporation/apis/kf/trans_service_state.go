package kf

import (
	"bytes"
	"encoding/json"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/type/type_kf"
)

const (
	apiTransServiceState = "/cgi-bin/kf/service_state/trans"
)

func TransServiceState(ctx *corporation.App, openKfid string, externalUserId string, serviceState int64, serviceUserId string) (*type_kf.TransferStateResp, error) {
	params := map[string]interface{}{
		"open_kfid":       openKfid,
		"external_userid": externalUserId,
		"service_state":   serviceState,
		"servicer_userid": serviceUserId,
	}

	payload, _ := json.Marshal(params)
	data, err := ctx.Client.HTTPPost(apiTransServiceState, bytes.NewReader(payload), "application/json;charset=utf-8")
	if err != nil {
		return nil, err
	}

	rsp := type_kf.TransferStateResp{}
	err = json.Unmarshal(data, &rsp)
	return &rsp, err
}
