package simcount

type CheckResp struct {
	ReqID    string    `json:"reqId"`
	Cid      string    `json:"cid"`
	Resource *Resource `json:"resource,omitempty"`
}

type Resource struct {
	SimCount int `json:"simCount,omitempty"`
}
