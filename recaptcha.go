package recaptcha

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const verifyURL = "https://www.google.com/recaptcha/api/siteverify"

type _Resp struct {
	Success       bool      `json:"success,omitempty"`
	Score         float64   `json:"score,omitempty"`
	Action        string    `json:"action,omitempty"`
	ChallengeTime time.Time `json:"challenge_ts,omitempty"`
	HostName      string    `json:"host,omitempty"`
	Errors        []string  `json:"error-codes,omitempty"`
}

// Recaptcha is a structure to handle recaptcha
type Recaptcha struct {
	secKey string
	client *http.Client
}

// Check whether the response is from
// human (returns true) or not (returns false).
func (r Recaptcha) Check(remoteIP, response string) (res bool, err error) {
	raw, err := r.client.PostForm(verifyURL, url.Values{
		"secret":   {r.secKey},
		"remoteip": {remoteIP},
		"response": {response},
	})
	if err != nil {
		return
	}
	defer raw.Body.Close()
	decoder := json.NewDecoder(raw.Body)
	var resp _Resp
	err = decoder.Decode(&resp)
	if err != nil {
		return
	}
	res = resp.Success
	return
}

// New returns an instance of Recaptcha.
func New(secretKey string) Recaptcha {
	return Recaptcha{secKey: secretKey, client: &http.Client{}}
}
