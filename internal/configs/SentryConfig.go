package configs

type SentryConfig struct {
	Dsn     string
	Release string
	IsDebug bool
}

func NewSentryConfig() *SentryConfig {
	return &SentryConfig{}
}
