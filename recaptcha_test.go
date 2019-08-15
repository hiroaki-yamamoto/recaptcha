package recaptcha_test

import (
	"reflect"
	test "testing"

	"github.com/hiroaki-yamamoto/recaptcha"
	"gotest.tools/assert"
)

func TestRecaptchaInit(t *test.T) {
	r := recaptcha.New("6LeIxAcTAAAAAGG-vFI1TnRWxMZNFuojJ4WifJWe")
	assert.Equal(t, reflect.TypeOf(r), reflect.TypeOf(recaptcha.Recaptcha{}))
}

func TestSuccess(t *test.T) {
	r := recaptcha.New("6LeIxAcTAAAAAGG-vFI1TnRWxMZNFuojJ4WifJWe")
	assert.Assert(t, r.Check("[::1]", "test_response"))
}
