package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/big"
	"time"

	"github.com/sayhellolexa/url-short/internal/domain"
	"github.com/sayhellolexa/url-short/internal/repository/inmemory"
)

// shortUrlLength - длина сокращенного URL
// alphabet - алфавит, из которого будет состоять сокращенный URL
const (
	alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

type UrlService struct {
	repo domain.UrlRepository
}

func NewUrlService() *UrlService {
	return &UrlService{repo: inmemory.NewUrlRepository()}
}

func (u *UrlService) ShortenUrl(longUrl string) (string, error) {
	if longUrl == "" {
		return "", errors.New("longUrl cannot be empty")
	}

	shortUrl := generateShortUrl(longUrl)
	err := u.repo.Save(longUrl, shortUrl)
	if err != nil {
		return "", err
	}

	return shortUrl, nil
}

func (u *UrlService) GetLongUrl(shortUrl string) (string, error) {
	longUrl, err := u.repo.GetLongUrl(shortUrl)
	if err != nil {
		return "", err
	}

	return longUrl, nil
}

/*
GenerateShortUrl - генерирует сокращенный URL
1. input - строка вида: longUrl + текущее время, время для того, чтобы одинаковые URL, сокращаемые в разное время
давали разные хеши (так называемая соль)
2. hash - вычисляется хеш
3. hashStr - преобразуем хеш в хекс-строку
4. Интерпретируем хекс-строку как число в 16-ричной системе
5. Делим число на 58 и берем остатки, каждый остаток - символ из алфавита
6. Берем первые 8 символов
*/
func generateShortUrl(longUrl string) string {
	input := longUrl + time.Now().String()
	hash := sha256.Sum256([]byte(input))
	hashStr := hex.EncodeToString(hash[:])
	bigInt := new(big.Int)
	bigInt.SetString(hashStr, 16)

	var result []byte
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := new(big.Int)

	for bigInt.Cmp(zero) > 0 {
		bigInt.DivMod(bigInt, base, mod)
		result = append(result, alphabet[mod.Int64()])
	}

	if len(result) > 8 {
		return string(result[:8])
	}
	return string(result)
}
