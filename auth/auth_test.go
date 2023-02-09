package auth

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Auth", func() {
	var (
		mockCtrl   *gomock.Controller
		httpClient *MockHTTPClient
		authClient *Client
		apiUrl     = "http://testing.url"
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		httpClient = NewMockHTTPClient(mockCtrl)
		authClient = NewClient(httpClient, apiUrl, "CLIENT_ID", "CLIENT_SECRET")
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("buildHeaders", func() {
		It("builds the headers", func() {
			headers := authClient.buildHeaders()
			Expect(headers["Accept"]).To(Equal("application/json"))
			Expect(headers["Content-Type"]).To(Equal("application/x-www-form-urlencoded"))
		})
	})

	Describe("buildParams", func() {
		It("builds the query params", func() {
			params := authClient.buildParams()
			Expect(params["grant_type"]).To(Equal("client_credentials"))
			Expect(params["client_id"]).To(Equal("CLIENT_ID"))
			Expect(params["client_secret"]).To(Equal("CLIENT_SECRET"))
		})
	})
})
