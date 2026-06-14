package profile

import "github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"

type QueryReq struct {
	Query string `json:"query"`
}

type ProfileResp struct {
	models.ReqInfoGetter
	Data        *ProfileData `json:"data,omitempty"`
	DataPresent bool         `json:"dataPresent"`
	Errors      []string     `json:"errors,omitempty"`
	Extensions  interface{}  `json:"extensions,omitempty"`
}

type ProfileData struct {
	Profile *Profile `json:"profile,omitempty"`
}

type Profile struct {
	Gender string `json:"gender,omitempty"`
	Age    string `json:"age,omitempty"`
}
