package linkService

import "context"

type Repo interface {
	SelectLink(ctx context.Context, abbreviatedLink string) (string, error)
	InsertLink(ctx context.Context, abbreviatedLink, originalLink string) error
}
