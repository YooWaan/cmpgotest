/*
 * main
 */

package main

import (
	"net/http"
	"context"

	"example/cmpgotest/api"
	"example/cmpgotest/infra"
)

func main() {
	_ = http.ListenAndServe(":8877", api.Handler(context.Background()))
}
