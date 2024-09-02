package kf

import (
	"bytes"
	"encoding/json"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/type/type_kf"
)

const (
	apiKfSyncMessage = "/cgi-bin/kf/sync_msg"
)

func SyncMessage(ctx *corporation.App, payload []byte) (*type_kf.SyncMsgSchema, error) {
	data, err := ctx.Client.HTTPPost(apiKfSyncMessage, bytes.NewReader(payload), "application/json;charset=utf-8")
	if err != nil {
		return nil, err
	}
	originInfo := type_kf.SyncMsgSchemaRaw{}
	if err = json.Unmarshal(data, &originInfo); err != nil {
		return nil, err
	}

	msgList := make([]*type_kf.KfMessage, 0)
	if len(originInfo.MsgList) > 0 {
		for _, msg := range originInfo.MsgList {
			newMsg := &type_kf.KfMessage{}
			if val, ok := msg["msgid"].(string); ok {
				newMsg.MsgID = val
			}
			if val, ok := msg["open_kfid"].(string); ok {
				newMsg.OpenKFID = val
			}
			if val, ok := msg["external_userid"].(string); ok {
				newMsg.ExternalUserID = val
			}
			if val, ok := msg["send_time"].(float64); ok {
				newMsg.SendTime = uint64(val)
			}
			if val, ok := msg["origin"].(float64); ok {
				newMsg.Origin = uint32(val)
			}

			if val, ok := msg["msgtype"].(string); ok {
				newMsg.MsgType = val
			}
			if newMsg.MsgType == "event" {
				if event, ok := msg["event"].(map[string]interface{}); ok {
					if eType, ok := event["event_type"].(string); ok {
						newMsg.EventType = eType
					}
				}
			}
			originData, err := json.Marshal(msg)
			if err != nil {
				return nil, err
			}
			newMsg.OriginData = originData
			msgList = append(msgList, newMsg)
		}
	}
	return &type_kf.SyncMsgSchema{
		ErrCode:    originInfo.ErrCode,
		ErrMsg:     originInfo.ErrMsg,
		NextCursor: originInfo.NextCursor,
		HasMore:    originInfo.HasMore,
		MsgList:    msgList,
	}, nil
}
