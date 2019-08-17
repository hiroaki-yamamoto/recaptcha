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
) (bool, error) {
	mock, close, err := stubs.CreateClientStub(handler)
	if err != nil {
		return false, err
	}
	defer close()
	r := New("6LeIxAcTAAAAAGG-vFI1TnRWxMZNFuojJ4WifJWe")
	r.client = mock
	return r.Check("[::1]", "test_response")
}

func TestSuccess(t *test.T) {
	result, err := performAccess(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			encoder := json.NewEncoder(w)
			err := encoder.Encode(_Resp{
				Success:  true,
				HostName: "localhost",
			})
			assert.NilError(t, err)
		},
	))
	assert.NilError(t, err)
	assert.Assert(t, result)
}

func TestFailure(t *test.T) {
	result, err := performAccess(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			encoder := json.NewEncoder(w)
			err := encoder.Encode(_Resp{
				Success:  false,
				HostName: "localhost",
			})
			assert.NilError(t, err)
		},
	))
	assert.NilError(t, err)
	assert.Assert(t, !result)
}
