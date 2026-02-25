package repository

import "context"

type URLRepository interface {
	InsertURL(ctx context.Context, longURL string) (int64, error)
	UpdateShortCode(ctx context.Context, id int64, code string) error
	GetURL(ctx context.Context, code string) (string, error)
	GetCodeByHash(ctx context.Context, longURL string) (string, error)
    SaveClick(ctx context.Context, code, ip, ua string) error
}

type CacheRepository interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

type AnalyticsRepository interface {
	PublishClick(ctx context.Context, code string) error
}
