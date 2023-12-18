package app

type App struct {
	c         Config
	shortener *URLShortener
}

func NewApp(c Config) *App {
	return &App{
		c: c,
	}
}
