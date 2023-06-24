package linkService

import "context"

type Repo interface {
	SelectLink(ctx context.Context, abbreviatedLink string) (string, error)
	InsertLink(ctx context.Context, abbreviatedLink, originalLink string) error
}
type Link interface {
	GetOriginalByAbbreviated(ctx context.Context, link string) (string, error)
	SaveOriginalLink(ctx context.Context, link string) (string, error)
}
