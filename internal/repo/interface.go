package repo

import "context"

type DatabaseOperator interface {
	Read(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Write(ctx context.Context, query string, args ...interface{}) (int64, error)
}
