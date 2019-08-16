package stubs

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
)

// CreateClientStub creates a stub to handle the request.
func CreateClientStub(hd http.Handler) (*http.Client, func()) {
	s := httptest.NewTLSServer(hd)
	s.Certificate().DNSNames = append(s.Certificate().DNSNames, "www.google.com")

	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(
				_ context.Context,
				network, _ string,
			) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
		},
	}
	return cli, s.Close
}
