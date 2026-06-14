package models

type RateLimit struct {
	Limit     int
	PeriodSec int
	Remaining int
	Reset     int
}

type Tariffication struct {
	Units int
}

type ReqInfo struct {
	RateLimit
	Tariffication
	RequestID string
}

func (r ReqInfo) GetReqInfo() ReqInfo {
	return r
}

type ReqInfoGetter interface {
	GetReqInfo() ReqInfo
}
