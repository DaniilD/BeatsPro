package configs

var configLocator *ConfigLocator

type ConfigLocator struct {
	serverConfig *ServerConfig
	dbConfig	 *DBConfig
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

func (configLocator *ConfigLocator) DBConfigInstance() *DBConfig{
	if configLocator.dbConfig == nil {
		configLocator.dbConfig = NewDBConfig()
	}

	return configLocator.dbConfig
}
