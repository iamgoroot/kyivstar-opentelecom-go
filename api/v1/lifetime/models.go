package lifetime

type CheckResp struct {
	ReqID    string    `json:"reqId"`
	Cid      string    `json:"cid"`
	Resource *Resource `json:"resource,omitempty"`
}

type Resource struct {
	LifetimeDuration *LifetimeDuration `json:"lifetimeDuration,omitempty"`
}

type LifetimeDuration struct {
	From     int    `json:"from,omitempty"`
	To       *int   `json:"to,omitempty"`
	TimeUnit string `json:"timeUnit,omitempty"`
}
