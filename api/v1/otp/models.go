package otp

type SendReq struct {
	To         string `json:"to"`
	TemplateID *int   `json:"templateId,omitempty"`
}

type SendResp struct {
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
	ReqID    string    `json:"reqId"`
	Cid      string    `json:"cid"`
	Resource *CheckRes `json:"resource,omitempty"`
}

type CheckRes struct {
	Status string `json:"status"`
}
