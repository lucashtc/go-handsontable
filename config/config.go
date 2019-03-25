package config

// DBConfig set connection database
type DBConfig struct {
	Drive    string
	User     string
	Password string
	Charset  string
	Host     string
	Port     string
}

// NewConfig ...
func NewConfig() *DBConfig {
	return &DBConfig{
		Drive:   "mysql",
		User:    "root",
		Host:    "127.0.0.1",
		Charset: "utf8",
	}
}
