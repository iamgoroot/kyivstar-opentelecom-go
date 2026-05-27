package flashcall

type CreateReq struct {
	To string `json:"to"`
}

type CreateResp struct {
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
	ReqID    string    `json:"reqId"`
	Cid      string    `json:"cid"`
	Resource *CheckRes `json:"resource,omitempty"`
}

type CheckRes struct {
	Status string `json:"status"`
}
