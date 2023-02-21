package kf

import (
	"bytes"
	"encoding/json"
	"github.com/fastwego/wxwork/kf"
	"github.com/fastwego/wxwork/kf/type_kf"
)

const (
	apiKfCustomerBatchGet = "/cgi-bin/kf/customer/batchget"
)

func KfCustomerBatchGet(ctx *kf.KfApp, payload []byte) (*type_kf.CustomerBatchGetSchema, error) {
	data, err := ctx.Client.HTTPPost(apiKfCustomerBatchGet, bytes.NewReader(payload), "application/json;charset=utf-8")
	if err != nil {
		return nil, err
	}

	rsp := type_kf.CustomerBatchGetSchema{}
	err = json.Unmarshal(data, &rsp)
	return &rsp, err
}
