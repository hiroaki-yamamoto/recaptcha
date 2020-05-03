package recaptcha_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hiroaki-yamamoto/recaptcha"
	"github.com/hiroaki-yamamoto/recaptcha/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var httpCli *mocks.MockIHttpClient
var rootCtrl *gomock.Controller
var recap *recaptcha.Recaptcha

func TestRecaptcha(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Recaptcha Suite")
}

var _ = BeforeEach(func() {
	rootCtrl = gomock.NewController(GinkgoT())
	httpCli = mocks.NewMockIHttpClient(rootCtrl)
	recap = recaptcha.New("6LeIxAcTAAAAAGG-vFI1TnRWxMZNFuojJ4WifJWe")
	recap.Client = httpCli
})

var _ = AfterEach(func() {
	rootCtrl.Finish()
})
