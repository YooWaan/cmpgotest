package plain

import (
	"context"
	"net/http"
	"strings"

	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"

	"example/cmpgotest/api"
	"example/cmpgotest/app"
)

func TestSuite(t *testing.T) {

	hd, ctx := api.HandlerWithContext(context.Background())
	server := httptest.NewServer(hd)
	if len(server.URL) == 0 {
		t.Errorf("failed: create server")
	}
	defer server.Close()
	serverURL := server.URL

	t.Run("Suite", func(t *testing.T) {

		t.Run("logic test ===> unittest", func(t *testing.T) {
			if app.Hello(ctx) != "Hello" {
				t.Errorf("Hello not match")
			}
		})

		t.Run("endpoint test ===> integration test", func(t *testing.T) {
			t.Run("GET", func(t *testing.T) {
				req := httptest.NewRequest(http.MethodGet, serverURL+"/", nil)
				w := httptest.NewRecorder()
				hd.ServeHTTP(w, req)

				res := w.Result()
				defer func() { _ = res.Body.Close() }()
				httpexpect.NewResponse(t, res).
					Status(http.StatusOK).
					Text().Equal("Hello")
			})
			t.Run("POST", func(t *testing.T) {
				req := httptest.NewRequest(http.MethodPost, serverURL + "/", strings.NewReader("{}"))
				w := httptest.NewRecorder()
				hd.ServeHTTP(w, req)

				res := w.Result()
				defer func() { _ = res.Body.Close() }()
				httpexpect.NewResponse(t, res).
					Status(http.StatusOK).
					JSON().
					Object().ContainsKey("value").ValueEqual("value", "Hello")
			})
		})

		t.Run("api test ===> e2e test", func(t *testing.T) {
			t.Run("GET", func(t *testing.T) {
				e := httpexpect.New(t, serverURL)
				e.GET("/").
					Expect().
					Status(http.StatusOK).
					Text().Equal("Hello")
			})
			t.Run("POST", func(t *testing.T) {
				e := httpexpect.New(t, serverURL+"/")
				e.POST("/").
					Expect().
					Status(http.StatusOK).
					JSON().
					Object().ContainsKey("value").ValueEqual("value", "Hello")
			})
		})
	})
}
