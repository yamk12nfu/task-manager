package usecases

import "context"

type Transaction interface {
	Do(context.Context, func(context.Context) error) error
}
