package sms

import "time"

type SendReq struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

type SendResp struct {
	ReqID string `json:"reqID"`
	MsgID string `json:"msgID"`
}

type CheckResp struct {
	ReqID  string     `json:"reqID"`
	MsgID  string     `json:"msgID"`
	Status string     `json:"status"`
	Date   *time.Time `json:"date"`
}
