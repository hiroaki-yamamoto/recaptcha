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

func TestSuccess(t *test.T) {
	mock, close, err := stubs.CreateClientStub(http.HandlerFunc(
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
	defer close()
	r := New("6LeIxAcTAAAAAGG-vFI1TnRWxMZNFuojJ4WifJWe")
	r.client = mock
	result, err := r.Check("[::1]", "test_response")
	assert.NilError(t, err)
	assert.Assert(t, result)
}
