package config

import "fmt"

type DBConfig struct {
	host     string
	port     int
	user     string
	password string
	name     string
	sslmode  string
}

func (dbConfig *DBConfig) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s "+
		"dbname=%s sslmode=%s", dbConfig.host, dbConfig.port, dbConfig.user,
		dbConfig.password, dbConfig.name, dbConfig.sslmode)
}
