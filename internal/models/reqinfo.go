package models

type RateLimit struct {
	Limit     int `json:"limit"`
	PeriodSec int `json:"periodSec"`
	Remaining int `json:"remaining"`
	Reset     int `json:"reset"`
}

type Tariffication struct {
	Units int `json:"units"`
}

type ReqInfo struct {
	RateLimit
	Tariffication
	RequestID string `json:"requestId"`
}

func (r ReqInfo) GetReqInfo() ReqInfo {
	return r
}

type ReqInfoGetter interface {
	GetReqInfo() ReqInfo
}
