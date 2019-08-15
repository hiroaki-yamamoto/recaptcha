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
