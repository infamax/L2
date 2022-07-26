package config

import "gopkg.in/yaml.v3"

type Config struct {
	DBConfig
	Port int
}

type file struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Sslmode  string `yaml:"sslmode"`
	} `yaml:"database"`
	Port int `yaml:"port"`
}

func ParseConfig(fileBytes []byte) (*Config, error) {
	cf := file{}
	err := yaml.Unmarshal(fileBytes, &cf)
	if err != nil {
		return nil, err
	}

	return &Config{
		DBConfig: DBConfig{
			host:     cf.Database.Host,
			port:     cf.Database.Port,
			user:     cf.Database.User,
			password: cf.Database.Password,
			name:     cf.Database.Name,
			sslmode:  cf.Database.Sslmode,
		},
		Port: cf.Port,
	}, nil
}
