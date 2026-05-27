package scoring

type CheckResp struct {
	ReqID    string    `json:"reqId"`
	Cid      string    `json:"cid"`
	Resource *Resource `json:"resource,omitempty"`
}

type Resource struct {
	ScoreBal float64 `json:"scoreBal,omitempty"`
}
