package ginkgoex

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"example/cmpgotest/api"
)

var (
	server    *httptest.Server
	serverURL string
	testContext context.Context
	testHandler http.Handler
)

func TestGinkotest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkotest Suite")
}

var _ = BeforeSuite(func() {
	hd, ctx := api.HandlerWithContext(context.Background())
	testContext = ctx
	testHandler = hd
	server = httptest.NewServer(hd)
	Expect(len(server.URL)).To(BeNumerically(">", 0))
	serverURL = server.URL
	log.Printf("server: %v", server.URL)
})

var _ = AfterSuite(func() {
	server.Close()
})
