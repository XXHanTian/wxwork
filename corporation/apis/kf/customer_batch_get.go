package kf

import (
	"bytes"
	"encoding/json"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/type/type_kf"
)

const (
	apiKfCustomerBatchGet = "/cgi-bin/kf/customer/batchget"
)

func KfCustomerBatchGet(ctx *corporation.App, payload []byte) (*type_kf.CustomerBatchGetSchema, error) {
	data, err := ctx.Client.HTTPPost(apiKfCustomerBatchGet, bytes.NewReader(payload), "application/json;charset=utf-8")
	if err != nil {
		return nil, err
	}

	rsp := type_kf.CustomerBatchGetSchema{}
	err = json.Unmarshal(data, &rsp)
	return &rsp, err
}
