package recaptcha // Needs to mock http request.

import (
	"encoding/json"
	"net/http"
	"reflect"
	test "testing"
)

import (
	"github.com/hiroaki-yamamoto/recaptcha/stubs"
	"gotest.tools/assert"
)

func TestRecaptchaInit(t *test.T) {
	r := New("6LeIxAcTAAAAAGG-vFI1TnRWxMZNFuojJ4WifJWe")
	assert.Equal(t, reflect.TypeOf(r), reflect.TypeOf(Recaptcha{}))
}

func performAccess(
	handler http.HandlerFunc,
) (Response, error) {
	mock, close, err := stubs.CreateClientStub(handler)
	if err != nil {
		return Response{}, err
	}
	defer close()
	r := New("6LeIxAcTAAAAAGG-vFI1TnRWxMZNFuojJ4WifJWe")
	r.client = mock
	return r.Check("[::1]", "test_response")
}

func TestSuccess(t *test.T) {
	rsp := Response{
		Success:  true,
		HostName: "localhost",
	}
	result, err := performAccess(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			encoder := json.NewEncoder(w)
			err := encoder.Encode(rsp)
			assert.NilError(t, err)
		},
	))
	assert.NilError(t, err)
	assert.DeepEqual(t, result, rsp)
}

func TestFailure(t *test.T) {
	rsp := Response{
		Success:  false,
		HostName: "localhost",
	}
	result, err := performAccess(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			encoder := json.NewEncoder(w)
			err := encoder.Encode(rsp)
			assert.NilError(t, err)
		},
	))
	assert.NilError(t, err)
	assert.DeepEqual(t, result, rsp)
}
