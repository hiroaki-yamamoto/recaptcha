package recaptcha_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/golang/mock/gomock"
	. "github.com/hiroaki-yamamoto/recaptcha"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const verifyURL = "https://www.google.com/recaptcha/api/siteverify"

var _ = Describe("Recaptcha", func() {
	Context("In Mocked HTTP Client", func() {
		var postCall *gomock.Call
		postCallFunc := func(res *Response, statusCode int) func(
			url string, param url.Values,
		) (*http.Response, error) {
			return func(
				url string, param url.Values,
			) (*http.Response, error) {
				recorder := httptest.NewRecorder()
				recorder.WriteHeader(statusCode)
				encoder := json.NewEncoder(recorder)
				Expect(encoder.Encode(res)).To(Succeed())
				return recorder.Result(), nil
			}
		}

		BeforeEach(func() {
			postCall = httpCli.EXPECT().PostForm(
				gomock.Eq(verifyURL),
				gomock.Eq(url.Values{
					"secret":   {recap.SecKey},
					"remoteip": {"localhost"},
					"response": {"test"},
				}),
			).Times(1)
		})

		Context("With Success Result", func() {
			var res *Response
			BeforeEach(func() {
				res = &Response{
					Success:  true,
					HostName: "localhost",
				}
				postCall = postCall.DoAndReturn(postCallFunc(res, http.StatusOK))
			})
			It("Should return the successful response.", func() {
				resp, err := recap.Check("localhost", "test")
				Expect(err).To(Succeed())
				Expect(resp).To(Equal(res))
			})
		})

		Context("With Failure Result", func() {
			var res *Response
			BeforeEach(func() {
				res = &Response{
					Success:  false,
					HostName: "localhost",
				}
				postCall = postCall.DoAndReturn(postCallFunc(res, http.StatusOK))
			})
			It("Should return the failure response.", func() {
				resp, err := recap.Check("localhost", "test")
				Expect(err).To(Succeed())
				Expect(resp).To(Equal(res))
			})
		})

		Context("With Invalid Payload", func() {
			var code int
			var resp *http.Response
			BeforeEach(func() {
				code = http.StatusOK
				rec := httptest.NewRecorder()
				rec.WriteHeader(code)
				rec.Write([]byte("Test Payload"))
				resp = rec.Result()
				postCall = postCall.DoAndReturn(func(
					url string, param url.Values,
				) (*http.Response, error) {
					return resp, nil
				})
			})
			It("Should raise an error.", func() {
				acResp, err := recap.Check("localhost", "test")
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal(
					"invalid character 'T' looking for beginning of value",
				))
				Expect(acResp).To(BeNil())
			})
		})

		Context("With Server Error", func() {
			var code int
			var resp *http.Response
			BeforeEach(func() {
				code = http.StatusInternalServerError
				rec := httptest.NewRecorder()
				rec.WriteHeader(code)
				rec.Write([]byte("Test Error"))
				resp = rec.Result()
				postCall = postCall.DoAndReturn(func(
					url string, param url.Values,
				) (*http.Response, error) {
					return resp, nil
				})
			})
			It("Should raise an error.", func() {
				acResp, err := recap.Check("localhost", "test")
				Expect(err).To(MatchError(
					&ResponseError{Response: resp, Body: "Test Error"},
				))
				Expect(err.Error()).To(Equal(
					fmt.Sprintf("Post %s: Returned %d: Test Error", verifyURL, code),
				))
				Expect(acResp).To(BeNil())
			})
		})

		Context("When PostForm raised an error", func() {
			BeforeEach(func() {
				postCall = postCall.Return(nil, errors.New("Test Error"))
			})
			It("Should raise an error", func() {
				acResp, err := recap.Check("localhost", "test")
				Expect(err).To(MatchError(errors.New("Test Error")))
				Expect(acResp).To(BeNil())
			})
		})
	})
})
