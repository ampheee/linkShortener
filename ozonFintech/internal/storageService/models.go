package storage

import "context"

type Storage interface {
	InsertLink(ctx context.Context)
	GetLink(ctx context.Context) error
}
