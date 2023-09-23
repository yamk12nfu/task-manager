package usecases

import "context"

type UniqueIDRepository interface {
	Issue(context.Context) (string, error)
}
