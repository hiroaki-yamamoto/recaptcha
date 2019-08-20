package recaptcha

import (
	"bytes"
	"encoding/json"
	"fmt"
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

// ResponseError represents the error of response from the server.
type ResponseError struct {
	Response *http.Response // Note: Response.Body is always closed in this case.
	Body     string
}

func (r ResponseError) Error() string {
	return fmt.Sprintf(
		"Post %s: Returned %d: %s", verifyURL, r.Response.StatusCode, r.Body,
	)
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
	if raw.StatusCode != http.StatusOK {
		var buf bytes.Buffer
		_, rerr := buf.ReadFrom(raw.Body)
		err = rerr
		if err == nil {
			err = ResponseError{Response: raw, Body: string(buf.String())}
		}
		return
	}
	decoder := json.NewDecoder(raw.Body)
	err = decoder.Decode(&res)
	return
}

// New returns an instance of Recaptcha.
func New(secretKey string) Recaptcha {
	return Recaptcha{secKey: secretKey, client: &http.Client{}}
}
