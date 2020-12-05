/*
 * infrastructure code
 */

package infra

import (
	"context"

	"example/cmpgotest/domain"
	"example/cmpgotest/app"
)

func ConstRepo() string { return domain.Hello }

func Inject(c context.Context) context.Context {
	return context.WithValue(c, app.KeyContext, app.Context{Repo:ConstRepo})
}
