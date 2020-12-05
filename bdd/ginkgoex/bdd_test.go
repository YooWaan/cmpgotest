package ginkgoex

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"example/cmpgotest/app"
)

var _ = Describe("BDD", func() {

	Describe("BDD Test", func() {
		Context("logic test ===> unit test", func() {
			It("Hello usecase", func() {
				Expect(app.Hello(testContext)).To(Equal("Hello"))
			})
		})

		Context("endpoint test ===> integration test", func() {
			It("GET", func() {
				req := httptest.NewRequest(http.MethodGet, serverURL + "/", nil)
				w := httptest.NewRecorder()
				testHandler.ServeHTTP(w, req)

				res := w.Result()
				Expect(res.StatusCode).Should(Equal(http.StatusOK))
				body, err := ioutil.ReadAll(res.Body)
				Expect(err).ShouldNot(HaveOccurred())
				_ = res.Body.Close()
				Expect(string(body)).To(Equal("Hello"))
			})

			It("POST", func() {
				req := httptest.NewRequest(http.MethodPost, serverURL + "/", strings.NewReader("{}"))
				w := httptest.NewRecorder()
				testHandler.ServeHTTP(w, req)

				res := w.Result()
				Expect(res.StatusCode).Should(Equal(http.StatusOK))
				body, err := ioutil.ReadAll(res.Body)
				Expect(err).ShouldNot(HaveOccurred())
				_ = res.Body.Close()

				log.Printf("Ret:%s", string(body))
				Expect(string(body)).To(Equal(`{"value":"Hello"}`))
			})
		})

		Context("api test ===> e2e test", func() {
			It("GET", func() {
				res, err := http.Get(serverURL + "/")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(res.StatusCode).Should(Equal(http.StatusOK))

				body, err := ioutil.ReadAll(res.Body)
				Expect(err).ShouldNot(HaveOccurred())
				_ = res.Body.Close()
				Expect(string(body)).To(Equal("Hello"))
			})
			It("POST", func() {
				res, err := http.Post(serverURL+"/", "application/json", strings.NewReader("{}"))
				Expect(err).ShouldNot(HaveOccurred())
				Expect(res.StatusCode).Should(Equal(http.StatusOK))

				body, err := ioutil.ReadAll(res.Body)
				Expect(err).ShouldNot(HaveOccurred())
				_ = res.Body.Close()

				log.Printf("Ret:%s", string(body))
				Expect(string(body)).To(Equal(`{"value":"Hello"}`))
			})
		})
	})
})
