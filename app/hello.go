/*
 * application code
 */

package app

import (
	"context"

	"example/cmpgotest/domain"
)

var KeyContext = struct{C string}{C:"context"}

type Context struct {
	domain.Repo
}

func Hello(c context.Context) string {
	actx := c.Value(KeyContext).(Context)
	return actx.Repo()
}
