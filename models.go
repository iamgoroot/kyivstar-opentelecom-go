package client

import (
	"time"
)

// Sms

type SmsSendReq struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

type SmsSendResp struct {
	ReqId string `json:"reqId"`
	MsgId string `json:"msgId"`
}

type SmsCheckResp struct {
	ReqId  string     `json:"reqId"`
	MsgId  string     `json:"msgId"`
	Status string     `json:"status"`
	Date   *time.Time `json:"date"`
}

// Verify sim

type VerifySimReq struct {
	ActivationHours int `json:"activationHours"`
}

type VerifySimResp struct {
	SimChanged int `json:"simChanged"`
	IsActive   int `json:"isActive"`
}

//Scoring

type ScoringResp struct {
	Score int `json:"score"`
}
