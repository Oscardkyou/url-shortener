package main

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
)

type StorageInterface interface {
	Save(originalURL string, shortKey string,) error
}

type URLShortener struct {
	storage StorageInterface
}

func NewURLShortener(storage StorageInterface) *URLShortener {
	return &URLShortener{
		storage: storage,
	}
}

func (us *URLShortener) Shorten(originalURL string) string {
	shortKey := generateShortKey()
	us.storage.Save(originalURL, shortKey)
	return shortKey
}

func main() {
	short := NewURLShortener(nil)
	cfg := Config{}
	if cfg.UseDb {

	}

}


type RedisStorage struct {
	redis *redis.Client
}

// 1) Хранение данных о бизнес модели и книге 
book = Book{name: "sdsds", author: "Автор А.А."}
book.GetName() // Получение названия книги 
//2) Сохранение данных о кнгие в базу данных 
book.Save()

bookStorage = BookStorage{}
bookStorage.Save(book)

