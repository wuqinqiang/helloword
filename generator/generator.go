package generator

import "context"

type Generator interface {
	Generate(ctx context.Context, words []string) (phrase string, err error)
}
