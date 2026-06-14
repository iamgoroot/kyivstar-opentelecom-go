package scoring

import "github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"

type CheckResp struct {
	models.ReqInfoGetter
	ReqID    string    `json:"reqId"`
	Cid      string    `json:"cid"`
	Resource *Resource `json:"resource,omitempty"`
}

type Resource struct {
	ScoreBal float64 `json:"scoreBal,omitempty"`
}
