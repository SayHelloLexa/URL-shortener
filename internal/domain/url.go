package domain

type UrlRepository interface {
	Save(longUrl string, shortUrl string) error
	GetLongUrl(shortUrl string) (string, error)
}
