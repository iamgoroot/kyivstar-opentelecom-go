package flashcall

import "github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"

type CreateReq struct {
	To string `json:"to"`
}

type CreateResp struct {
	models.ReqInfoGetter
	Cid      string     `json:"cid"`
	ReqID    string     `json:"reqId"`
	Resource *CreateRes `json:"resource,omitempty"`
}

type CreateRes struct {
	Status string `json:"status"`
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
