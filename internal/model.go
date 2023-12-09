type URLShortener struct {
	urls map[string]string
}

func (us *URLShortener) Shorten(originalURL string) string {
	shortKey := generateShortKey()
	us.urls[shortKey] = originalURL
	return shortKey
}

func (us *URLShortener) GetOriginalURL(shortKey string) (string, bool) {
	originalURL, found := us.urls[shortKey]
	return originalURL, found
}
