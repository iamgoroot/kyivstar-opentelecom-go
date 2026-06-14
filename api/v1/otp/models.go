package otp

import "github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"

type SendReq struct {
	To         string `json:"to"`
	TemplateID *int   `json:"templateId,omitempty"`
}

type SendResp struct {
	models.ReqInfoGetter
	Cid      string   `json:"cid"`
	ReqID    string   `json:"reqId"`
	Resource *SendRes `json:"resource,omitempty"`
}

type SendRes struct {
	Status    string `json:"status"`
	MessageID string `json:"messageId,omitempty"`
}

type CheckReq struct {
	SubscriberID   string `json:"subscriberId"`
	ValidationCode string `json:"validationCode"`
}

type CheckResp struct {
	models.ReqInfoGetter
	ReqID    string    `json:"reqId"`
	Cid      string    `json:"cid"`
	Resource *CheckRes `json:"resource,omitempty"`
}

type CheckRes struct {
	Status string `json:"status"`
}
