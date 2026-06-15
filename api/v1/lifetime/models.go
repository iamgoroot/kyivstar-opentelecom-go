package lifetime

import "github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"

type CheckResp struct {
	models.ReqInfoGetter
	ReqID    string    `json:"reqId"`
	Cid      string    `json:"cid"`
	Resource *Resource `json:"resource,omitempty"`
}

type Resource struct {
	Duration *Duration `json:"lifetimeDuration,omitempty"`
}

type Duration struct {
	From     int    `json:"from,omitempty"`
	To       *int   `json:"to,omitempty"`
	TimeUnit string `json:"timeUnit,omitempty"`
}
