package rcs

type RcsTextReq struct {
	From               string      `json:"from"`
	To                 string      `json:"to"`
	MessageTtlSec      *int        `json:"messageTtlSec,omitempty"`
	CallbackNumber     *string     `json:"callbackNumber,omitempty"`
	ContentExtendedRcs ContentText `json:"contentExtendedRcs"`
}

type RcsSuggestionReq struct {
	From               string            `json:"from"`
	To                 string            `json:"to"`
	MessageTtlSec      *int              `json:"messageTtlSec,omitempty"`
	CallbackNumber     *string           `json:"callbackNumber,omitempty"`
	ContentExtendedRcs ContentSuggestion `json:"contentExtendedRcs"`
}

type RcsRichCardReq struct {
	From               string          `json:"from"`
	To                 string          `json:"to"`
	MessageTtlSec      *int            `json:"messageTtlSec,omitempty"`
	CallbackNumber     *string         `json:"callbackNumber,omitempty"`
	ContentExtendedRcs ContentRichCard `json:"contentExtendedRcs"`
}

type ContentText struct {
	Text string `json:"text"`
}

type ContentSuggestion struct {
	Text        string       `json:"text"`
	Suggestions []Suggestion `json:"suggestions,omitempty"`
}

type ContentRichCard struct {
	StandaloneCard *StandaloneCard `json:"standaloneCard,omitempty"`
}

type StandaloneCard struct {
	ThumbnailImageAlignment *string      `json:"thumbnailImageAlignment,omitempty"`
	CardOrientation         *string      `json:"cardOrientation,omitempty"`
	CardContent             *CardContent `json:"cardContent,omitempty"`
}

type CardContent struct {
	Title       string       `json:"title,omitempty"`
	Description string       `json:"description,omitempty"`
	Media       *Media       `json:"media,omitempty"`
	Suggestions []Suggestion `json:"suggestions,omitempty"`
}

type Media struct {
	Height       string `json:"height,omitempty"`
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"`
	FileUrl      string `json:"fileUrl,omitempty"`
}

type Suggestion struct {
	Type          string `json:"type,omitempty"`
	Text          string `json:"text,omitempty"`
	OpenUrlAction string `json:"openUrlAction,omitempty"`
}

type SendResp struct {
	ReqID string `json:"reqId"`
	MsgID string `json:"msgId"`
}

type CheckResp struct {
	ReqID  string `json:"reqId"`
	MsgID  string `json:"msgId"`
	Status string `json:"status"`
	Date   string `json:"date"`
}
