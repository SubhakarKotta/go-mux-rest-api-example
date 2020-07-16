package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Host     string
	Port     int
	Name     string
	Charset  string
}

func GetConfig() *Config {
	dbConfig := DBConfig{
		Dialect:  "mysql",
		Username: "mysqluser",
		Password: "mysqlpw",
		Port:     3306,
		Host:     "localhost",
		Name:     "eventuate",
		Charset:  "utf8",
	}
	config := Config{
		DB: &dbConfig,
	}
	return &config
}
