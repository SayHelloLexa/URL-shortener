package inmemory

import (
	"errors"
	"sync"

	"github.com/sayhellolexa/url-short/internal/domain"
	"github.com/sayhellolexa/url-short/internal/model"
)

type urlRepository struct {
	mu        sync.Mutex
	urls      []model.Url
	idCounter int
}

func NewUrlRepository() domain.UrlRepository {
	return &urlRepository{
		mu:        sync.Mutex{},
		urls:      make([]model.Url, 0),
		idCounter: 0,
	}
}

func (r *urlRepository) Save(longUrl string, shortUrl string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := r.idCounter
	r.idCounter++

	r.urls = append(r.urls, model.Url{
		Id:       id,
		LongUrl:  longUrl,
		ShortUrl: shortUrl,
	})

	return nil
}

func (r *urlRepository) GetLongUrl(shortUrl string) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if shortUrl == "" {
		return "", errors.New("empty short url")
	}

	if len(r.urls) == 0 {
		return "", errors.New("nil urls storage")
	}

	for _, url := range r.urls {
		if url.ShortUrl == shortUrl {
			return url.LongUrl, nil
		}
	}

	return "", errors.New("url not found")
}
