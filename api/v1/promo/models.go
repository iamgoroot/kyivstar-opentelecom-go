package promo

type CreateSMSReq struct {
	From         string `json:"from"`
	Text         string `json:"text"`
	CampaignType string `json:"campaignType"`
}

type CreateViberReq struct {
	From          string `json:"from"`
	Text          string `json:"text"`
	CampaignType  string `json:"campaignType"`
	StartDate     string `json:"startDate,omitempty"`
	Action        string `json:"action,omitempty"`
	Caption       string `json:"caption,omitempty"`
	MessageTtlSec *int   `json:"messageTtlSec,omitempty"`
}

type CreateRCSReq struct {
	From           string `json:"from"`
	Text           string `json:"text"`
	CampaignType   string `json:"campaignType"`
	StartDate      string `json:"startDate,omitempty"`
	Action         string `json:"action,omitempty"`
	Caption        string `json:"caption,omitempty"`
	Title          string `json:"title,omitempty"`
	RcsContentType string `json:"rcsContentType,omitempty"`
	MessageTtlSec  *int   `json:"messageTtlSec,omitempty"`
}

type Promo struct {
	ID                    string          `json:"id,omitempty"`
	AuthorUsername        string          `json:"authorUsername,omitempty"`
	Status                string          `json:"status,omitempty"`
	StartDate             string          `json:"startDate,omitempty"`
	EndDate               string          `json:"endDate,omitempty"`
	TextToSend            string          `json:"textToSend,omitempty"`
	TextUa                string          `json:"textUa,omitempty"`
	RcsContentType        string          `json:"rcsContentType,omitempty"`
	Title                 string          `json:"title,omitempty"`
	NextAvailableStatuses []string        `json:"nextAvailableStatuses,omitempty"`
	MessageContent        *MessageContent `json:"messageContent,omitempty"`
}

type MessageContent struct {
	MessageParamCount int         `json:"messageParamCount,omitempty"`
	SmsContent        *SMSContent `json:"smsContent,omitempty"`
}

type SMSContent struct {
	Text         string `json:"text,omitempty"`
	SourceNumber string `json:"sourceNumber,omitempty"`
	MessageTtl   int    `json:"messageTtl,omitempty"`
}

type ListResp struct {
	ReqID         string  `json:"reqId"`
	Promos        []Promo `json:"promos,omitempty"`
	TotalPages    int     `json:"totalPages,omitempty"`
	TotalElements int     `json:"totalElements,omitempty"`
	Number        int     `json:"number,omitempty"`
	Size          int     `json:"size,omitempty"`
}

type AddAudienceReq struct {
	Audience []AudienceMember `json:"audience"`
}

type AudienceMember struct {
	Params      []string `json:"params"`
	PhoneNumber string   `json:"phoneNumber"`
}

type AddAudienceResp struct {
	ReqID string `json:"reqId"`
	Name  string `json:"name"`
}

type AddImageResp struct {
	ReqID   string `json:"reqId"`
	Success bool   `json:"success"`
}

type PromoStat struct {
	SentCount                        int `json:"sentCount,omitempty"`
	DeliveriesCount                  int `json:"deliveriesCount,omitempty"`
	UnmatchedCount                   int `json:"unmatchedCount,omitempty"`
	DeliveriesPortionsCount          int `json:"deliveriesPortionsCount,omitempty"`
	DeliveriesUnmatchedPortionsCount int `json:"deliveriesUnmatchedPortionsCount,omitempty"`
	DeliveriesInternalPortionsCount  int `json:"deliveriesInternalPortionsCount,omitempty"`
	DeliveriesExternalPortionsCount  int `json:"deliveriesExternalPortionsCount,omitempty"`
	UndeliveredCount                 int `json:"undeliveredCount,omitempty"`
	UnknownStatusCount               int `json:"unknownStatusCount,omitempty"`
	CanceledByContactPolicyCount     int `json:"canceledByContactPolicyCount,omitempty"`
	SeenCount                        int `json:"seenCount,omitempty"`
	BlacklistedCount                 int `json:"blacklistedCount,omitempty"`
	DeclinedCount                    int `json:"declinedCount,omitempty"`
	ExpiredCount                     int `json:"expiredCount,omitempty"`
	WasNotSent                       int `json:"wasNotSent,omitempty"`
}
