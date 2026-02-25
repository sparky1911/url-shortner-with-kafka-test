package service

import (
	"context"

	"github.com/your-username/url-shortener/internal/repository"
	"github.com/your-username/url-shortener/pkg/utils"
)

type URLService struct {

	repo      repository.URLRepository
	cache     repository.CacheRepository
	analytics repository.KafkaRepository
}


func NewURLService(r repository.URLRepository, c repository.CacheRepository, a repository.KafkaRepository) *URLService {
	return &URLService{repo: r, cache: c, analytics: a}
}

func (s *URLService) ShortenUrl(ctx context.Context, longURL string) (string, error) {
	id, err := s.repo.InsertURL(ctx, longURL)
	if err != nil {
		return "", err
	}

	shortCode := utils.Encode(uint64(id))

	err = s.repo.UpdateShortCode(ctx, id, shortCode)
	if err != nil {
		return "", err
	}

	_ = s.cache.Set(ctx, shortCode, longURL)

	return shortCode, nil
}

func (s *URLService) GetOriginalURL(ctx context.Context, code string ,ip string, userAgent string ) (string, error) {

	longURL, err := s.cache.Get(ctx, code)

	if err != nil || longURL == "" {
		longURL, err = s.repo.GetURL(ctx, code)
		if err != nil {
			return "", err
		}
		_ = s.cache.Set(ctx, code, longURL)
	}
	if longURL != "" {
		go func() {
			_ = s.analytics.PublishClick(context.Background(), code,ip,userAgent)
		}()
	}

	return longURL, nil

}
