package shortener

func (us *URLShortener) Shorten(originalURL string) string {
	shortKey := generateShortKey()
	us.urls[shortKey] = originalURL
	return shortKey
}
