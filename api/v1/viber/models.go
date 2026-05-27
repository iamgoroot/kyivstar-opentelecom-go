package viber

type TransactionReq struct {
	From          string `json:"from"`
	To            string `json:"to"`
	Text          string `json:"text"`
	MessageTtlSec *int   `json:"messageTtlSec,omitempty"`
}

type PromotionTextReq struct {
	From          string `json:"from"`
	To            string `json:"to"`
	Text          string `json:"text"`
	MessageTtlSec *int   `json:"messageTtlSec,omitempty"`
}

type PromotionImageReq struct {
	From            string             `json:"from"`
	To              string             `json:"to"`
	ContentExtended ContentExtendedImg `json:"contentExtended"`
	MessageTtlSec   *int               `json:"messageTtlSec,omitempty"`
}

type PromotionActionReq struct {
	From            string                `json:"from"`
	To              string                `json:"to"`
	Text            string                `json:"text"`
	ContentExtended ContentExtendedAction `json:"contentExtended"`
	MessageTtlSec   *int                  `json:"messageTtlSec,omitempty"`
}

type ContentExtendedImg struct {
	Img string `json:"img"`
}

type ContentExtendedAction struct {
	Img     string `json:"img"`
	Caption string `json:"caption"`
	Action  string `json:"action"`
}

type SendResp struct {
	ReqID string `json:"reqId"`
	Mid   string `json:"mid"`
}

type CheckResp struct {
	ReqID  string `json:"reqId"`
	Mid    string `json:"mid"`
	Status string `json:"status"`
	Date   string `json:"date"`
}
