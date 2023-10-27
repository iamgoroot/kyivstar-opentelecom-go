package client

type Client interface {
	Sms
	VerifySim
	Scoring
}

type Sms interface {
	Send(req SmsSendReq) (SmsSendResp, error)
	Check(msgID string) (SmsCheckResp, error)
}
type VerifySim interface {
	VerifySim(phoneNumber string, params VerifySimReq) (VerifySimResp, error)
}

type Scoring interface {
	Scoring(phoneNumber string, modelId int) (ScoringResp, error)
}
