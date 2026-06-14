package devicecheck

import "github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"

type CheckResp struct {
	models.ReqInfoGetter
	ReqID   string   `json:"reqId"`
	Cid     string   `json:"cid"`
	ImeiRes string   `json:"imeiRes,omitempty"`
	ImeiCnt *float64 `json:"imeiCount,omitempty"`
}

type CheckRespWithResource struct {
	models.ReqInfoGetter
	ReqID    string    `json:"reqId"`
	Cid      string    `json:"cid"`
	Resource *Resource `json:"resource,omitempty"`
}

type Resource struct {
	ImeiRes   string   `json:"imeiRes,omitempty"`
	ImeiCount *float64 `json:"imeiCount,omitempty"`
}
