package recaptcha

import (
	"net/http"
	"net/url"
)

// IHttpClient abstracts http client (net/http) for testing.
type IHttpClient interface {
	PostForm(url string, data url.Values) (resp *http.Response, err error)
}

// IRecaptcha abstracts recaptcha struct for testing / using.
type IRecaptcha interface {
	Check(remoteIP, response string) (res *Response, err error)
}
