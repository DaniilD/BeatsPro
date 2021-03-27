package configs

var configLocator *ConfigLocator

type ConfigLocator struct {
	serverConfig *ServerConfig
	dbConfig     *DBConfig
	sentryConfig *SentryConfig
}

func GetConfigLocator() *ConfigLocator {
	if configLocator == nil {
		configLocator = &ConfigLocator{}
	}

	return configLocator
}

func (configLocator *ConfigLocator) ServerConfigInstance() *ServerConfig {
	if configLocator.serverConfig == nil {
		configLocator.serverConfig = NewServerConfig()
	}

	return configLocator.serverConfig
}

func (configLocator *ConfigLocator) DBConfigInstance() *DBConfig {
	if configLocator.dbConfig == nil {
		configLocator.dbConfig = NewDBConfig()
	}

	return configLocator.dbConfig
}

func (configLocator *ConfigLocator) SentryConfigInstance() *SentryConfig {
	if configLocator.sentryConfig == nil {
		configLocator.sentryConfig = NewSentryConfig()
	}

	return configLocator.sentryConfig
}
