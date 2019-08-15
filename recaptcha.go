package recaptcha

import "net/http"

const veryfyURL = "https://www.google.com/recaptcha/api/siteverify"

// Recaptcha is a structure to handle recaptcha
type Recaptcha struct {
	secKey string
	client *http.Client
}

// Check whether the response is from
// human (returns true) or not (returns false).
func (r Recaptcha) Check(remoteIP string, response string) (bool, err error) {

}

// New returns an instance of Recaptcha.
func New(secretKey string) Recaptcha {
	return Recaptcha{secKey: secretKey, client: &http.Client{}}
}

type _Resp struct {
	Success       bool     `json:"success,omitempty"`
	Score         float64  `json:"score,omitempty"`
	Action        string   `json:"action,omitempty"`
	ChallengeTime string   `json:"challenge_ts,omitempty"`
	HostName      string   `json:"host,omitempty"`
	Errors        []string `json:"error-codes,omitempty"`
}
