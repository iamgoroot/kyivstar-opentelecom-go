package multichannel

import "github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"

type SendReq struct {
	To           string        `json:"to"`
	SmsContent   *SmsContent   `json:"smsContent,omitempty"`
	ViberContent *ViberContent `json:"viberContent,omitempty"`
}

type SmsContent struct {
	Priority       int     `json:"priority"`
	From           string  `json:"from"`
	Text           string  `json:"text"`
	MessageTTLSec  int     `json:"messageTtlSec"`
	CallbackNumber *string `json:"callbackNumber,omitempty"`
}

type ViberContent struct {
	Priority      int     `json:"priority"`
	From          string  `json:"from"`
	PromoType     string  `json:"promoType,omitempty"`
	Text          string  `json:"text"`
	MessageTTLSec int     `json:"messageTtlSec"`
	Img           *string `json:"img,omitempty"`
	Caption       *string `json:"caption,omitempty"`
	Action        *string `json:"action,omitempty"`
}

type SendResp struct {
	models.ReqInfoGetter
	MultiMsgID string `json:"multiMsgId"`
}

type CheckResp struct {
	models.ReqInfoGetter
	Date       string    `json:"date"`
	MultiMsgID string    `json:"multiMsgId"`
	Status     string    `json:"status"`
	BearerType string    `json:"bearerType,omitempty"`
	Reports    []*Report `json:"reports,omitempty"`
}

type Report struct {
	BearerType string `json:"bearerType"`
	Date       string `json:"date"`
	Mid        string `json:"mid"`
	State      string `json:"state"`
}
