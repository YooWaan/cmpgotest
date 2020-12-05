package conveyex

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"

	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"example/cmpgotest/api"
	"example/cmpgotest/app"
)

func TestSute(t *testing.T) {

	hd, ctx := api.HandlerWithContext(context.Background())
	server := httptest.NewServer(hd)
	if len(server.URL) == 0 {
		t.Errorf("failed: create server")
	}
	defer server.Close()
	serverURL := server.URL

	readBody := func(res *http.Response) (string, error) {
		body, err := ioutil.ReadAll(res.Body)
		_ = res.Body.Close()
		if err != nil {
			return "", err
		}
		return string(body), nil
	}

	Convey("TestSuite", t, func() {

		Convey("logic test ===> usecase test", func() {
			Convey("Hello Usecase", func() {
				So(app.Hello(ctx), ShouldEqual, "Hello")
			})
		})

		Convey("endpoint test ===> integration test", func() {
			Convey("GET", func() {
				req := httptest.NewRequest(http.MethodGet, serverURL + "/", nil)
				w := httptest.NewRecorder()
				hd.ServeHTTP(w, req)

				res := w.Result()
				defer func() { _ = res.Body.Close() }()
				So(res.StatusCode, ShouldEqual, http.StatusOK)
				text, err := readBody(res)
				So(err, ShouldBeNil)
				So(text, ShouldEqual, "Hello")
			})

			Convey("POST", func() {
				req := httptest.NewRequest(http.MethodPost, serverURL + "/", strings.NewReader("{}"))
				w := httptest.NewRecorder()
				hd.ServeHTTP(w, req)

				res := w.Result()
				defer func() { _ = res.Body.Close() }()
				So(res.StatusCode, ShouldEqual, http.StatusOK)
				text, err := readBody(res)
				So(err, ShouldBeNil)
				So(text, ShouldEqual, `{"value":"Hello"}`)
			})
		})

		Convey("API Test", func() {

			Convey("GET", func() {
				res, err := http.Get(serverURL + "/")
				defer func() { _ = res.Body.Close() }()
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, http.StatusOK)

				text, err := readBody(res)
				So(err, ShouldBeNil)
				So(text, ShouldEqual, "Hello")
			})

			Convey("POST", func() {
				res, err := http.Post(serverURL+"/", "appliaction/json", strings.NewReader("{}"))
				defer func() { _ = res.Body.Close() }()
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, http.StatusOK)

				text, err := readBody(res)
				So(err, ShouldBeNil)
				So(text, ShouldEqual, `{"value":"Hello"}`)
			})

		})

	})
}
