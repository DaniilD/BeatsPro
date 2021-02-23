package configs

type DBConfig struct {
	Name     string
	Host     string
	User     string
	Password string
	Port     int
}


func NewDBConfig() *DBConfig {
	return &DBConfig{}
}
