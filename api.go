package ksopentelecom

import (
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/devicecheck"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/flashcall"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/lifetime"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/multichannel"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/otp"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/profile"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/promo"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/rcs"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/scoring"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/simcheck"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/simcount"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/sms"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/viber"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type (
	SMS          = sms.Service
	RCS          = rcs.Service
	Viber        = viber.Service
	Promo        = promo.Service
	Multichannel = multichannel.Service
	SimCheck     = simcheck.Service
	SimCount     = simcount.Service
	Scoring      = scoring.Service
	Lifetime     = lifetime.Service
	DeviceCheck  = devicecheck.Service
	OTP          = otp.Service
	FlashCall    = flashcall.Service
	Profile      = profile.Service
)

type V1Client struct {
	SMS sms.Service
	RCS
	Viber
	Promo
	Multichannel
	SimCheck
	SimCount
	Scoring
	Lifetime
	DeviceCheck
	OTP
	FlashCall
	Profile
}

func createV1Client(ksClient client.Client) (V1Client, error) {
	return V1Client{
		SMS:          sms.NewService(ksClient),
		RCS:          rcs.NewService(ksClient),
		Viber:        viber.NewService(ksClient),
		Promo:        promo.NewService(ksClient),
		Multichannel: multichannel.NewService(ksClient),
		SimCheck:     simcheck.NewService(ksClient),
		SimCount:     simcount.NewService(ksClient),
		Scoring:      scoring.NewService(ksClient),
		Lifetime:     lifetime.NewService(ksClient),
		DeviceCheck:  devicecheck.NewService(ksClient),
		OTP:          otp.NewService(ksClient),
		FlashCall:    flashcall.NewService(ksClient),
		Profile:      profile.NewService(ksClient),
	}, nil
}
