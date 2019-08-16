package stubs

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
)

type _RewriteTransport struct {
	Transport http.RoundTripper
	URL       *url.URL
}

func (t _RewriteTransport) RoundTrip(
	req *http.Request,
) (*http.Response, error) {
	// note that url.URL.ResolveReference doesn't work here
	// since t.u is an absolute url
	req.URL.Scheme = t.URL.Scheme
	req.URL.Host = t.URL.Host
	req.URL.Path = path.Join(t.URL.Path, req.URL.Path)
	rt := t.Transport
	if rt == nil {
		rt = http.DefaultTransport
	}
	return rt.RoundTrip(req)
}

// CreateClientStub creates a stub to handle the request.
func CreateClientStub(hd http.Handler) (
	cli *http.Client, close func(), err error,
) {
	s := httptest.NewServer(hd)
	u, err := url.Parse(s.URL)
	if err != nil {
		return
	}

	cli = &http.Client{Transport: _RewriteTransport{URL: u}}
	close = s.Close
	return
}
