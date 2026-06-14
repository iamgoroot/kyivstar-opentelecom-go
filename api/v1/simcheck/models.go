package simcheck

import "github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"

type CheckResp struct {
	models.ReqInfoGetter
	ReqID    string    `json:"reqId"`
	Cid      string    `json:"cid"`
	Resource *Resource `json:"resource,omitempty"`
}

type Resource struct {
	SimChanged     *bool `json:"simChanged,omitempty"`
	SimTypeChanged *bool `json:"simTypeChanged,omitempty"`
	GeoChanged     *bool `json:"geoChanged,omitempty"`
	ImeiChanged    *bool `json:"imeiChanged,omitempty"`
	CallForwarding *bool `json:"callForwarding,omitempty"`
}
