package recaptcha

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const verifyURL = "https://www.google.com/recaptcha/api/siteverify"

// Response represents the response from the server. It includes whethere
// the captcha is succeeded, the score (v3), action, time of the challenge,
// the name of the host, and error codes.
type Response struct {
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
func (r Recaptcha) Check(remoteIP, response string) (res Response, err error) {
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
	err = decoder.Decode(&res)
	return
}

// New returns an instance of Recaptcha.
func New(secretKey string) Recaptcha {
	return Recaptcha{secKey: secretKey, client: &http.Client{}}
}
