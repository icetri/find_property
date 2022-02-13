package config

type Config struct {
	HTTP        *HTTP     `yaml:"http"`
	PostgresDsn string    `yaml:"postgres_dsn"`
	Telegram    *Telegram `yaml:"telegram"`
}

type HTTP struct {
	Port string `yaml:"server_port"`
}

type Telegram struct {
	TelegramToken string `yaml:"telegram_token"`
	ChatID        string `yaml:"chat_id"`
}
