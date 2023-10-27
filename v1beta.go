package client

import (
	"fmt"
)

type v1beta struct {
	*requester
}

// Send Відправка SMS
func (c v1beta) Send(req SmsSendReq) (SmsSendResp, error) {
	return post[SmsSendReq, SmsSendResp](c.requester, "sms", req)
}

// Check Перевірка статусу SMS
func (c v1beta) Check(msgID string) (resp SmsCheckResp, err error) {
	url := fmt.Sprint("sms/", msgID)
	return get[SmsCheckResp](c.requester, url)
}

// VerifySim Перевірка заміни SIM карти
func (c v1beta) VerifySim(phoneNumber string, params VerifySimReq) (VerifySimResp, error) {
	url := fmt.Sprintf("subscribers/%s/verify-sim", phoneNumber)
	return post[VerifySimReq, VerifySimResp](c.requester, url, params)
}

// Scoring Фінансовий скоринг абонента
func (c v1beta) Scoring(phoneNumber string, modelId int) (ScoringResp, error) {
	if modelId < 0 {
		modelId = 7
	}
	url := fmt.Sprintf("subscribers/%s/score?modelId=%d", phoneNumber, modelId)
	return get[ScoringResp](c.requester, url)
}
