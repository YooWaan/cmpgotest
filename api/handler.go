/*
 * presentation code
 */

package api

import (
	"net/http"
	"context"

	"github.com/gin-gonic/gin"

	"example/cmpgotest/app"
	"example/cmpgotest/infra"
)

func HandlerWithContext(c context.Context) (http.HandlerFunc, context.Context) {
	ctx := infra.Inject(c)
	hd := Handle()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hd.ServeHTTP(w, r.WithContext(ctx))
	}), ctx
}

func Handler(c context.Context) http.HandlerFunc {
	hd, _ := HandlerWithContext(c)
	return hd
}

func Handle() *gin.Engine {
	e := gin.Default()
	e.GET("/", GETHello)
	e.POST("/", POSTHello)
	return e
}

func GETHello(c *gin.Context) {
	c.String(http.StatusOK, "%s", app.Hello(c.Request.Context()))
}

func POSTHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"value": app.Hello(c.Request.Context())})
}
