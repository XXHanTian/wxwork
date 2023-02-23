package kf

import (
	"bytes"
	"encoding/json"
	"github.com/fastwego/wxwork/kf"
)

const (
	apiKfAccountList = "/cgi-bin/kf/account/list"
)

func KfAccountList(ctx *kf.KfApp, offset, limit int64) ([]byte, error) {
	params := map[string]interface{}{
		"offset": offset,
		"limit":  limit,
	}

	payload, _ := json.Marshal(params)
	data, err := ctx.Client.HTTPPost(apiKfAccountList, bytes.NewReader(payload), "application/json;charset=utf-8")
	return data, err
}
