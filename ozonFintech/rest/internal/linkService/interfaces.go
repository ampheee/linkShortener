package linkService

import "context"

type Repo interface {
	SelectLink(ctx context.Context, originalLink string) error
	InsertLink(ctx context.Context, abbreviatedLink string) error
}
type Link interface {
	GetAbbreviatedLink(ctx context.Context, link string) (LinkDTO, error)
	SaveOriginalLink(ctx context.Context, link string) (LinkDTO, error)
}
