package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Host     string `yaml:"POSTGRES_HOST" env:"POSTGRES_HOST" env-default:"localhost"`
	Port     uint16 `yaml:"POSTGRES_PORT" env:"POSTGRES_PORT" env-default:"5432"`
	User     string `yaml:"POSTGRES_USER" env:"POSTGRES_USER" env-default:"user" `
	Pass     string `yaml:"POSTGRES_PASS" env:"POSTGRES_PASS" env-default:"secret"`
	Database string `yaml:"POSTGRES_DB" env:"POSTGRES_DB" env-default:"postgres_db"`
	MinCons  int32  `yaml:"POSTGRES_MIN_CON" env:"POSTGRES_MIN_CON" env-default:"5"`
	MaxCons  int32  `yaml:"POSTGRES_MAX_CON" env:"POSTGRES_MAX_CON" env-default:"10"`
}

// подгружаем данные ямла
func loadFromYAML(cfg *Config) error {
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return fmt.Errorf("failed to read yaml file: %w", err)
	}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return fmt.Errorf("failed to unmarshal yaml: %w", err)
	}
	return nil
}
func New() (*Config, error) {
	cfg := &Config{}
	if err := loadFromYAML(cfg); err != nil {
		fmt.Println("No YAML config loaded, falling back to defaults and env vars...")
	}
	return cfg, nil
}

func (c *Config) Allinfo() {
	fmt.Printf("Host: %s\n", c.Host)
	fmt.Printf("Port: %d\n", c.Port)
	fmt.Printf("User: %s\n", c.User)
	fmt.Printf("Pass: %s\n", c.Pass)
	fmt.Printf("Database: %s\n", c.Database)
	fmt.Printf("MinCons: %d\n", c.MinCons)
	fmt.Printf("MaxCons: %d\n", c.MaxCons)
}

// строка подлкючения пула конекшеннов
func (c *Config) DatabaseURL() string {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?pool_max_conns=%d&pool_min_conns=%d",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.Database,
		c.MaxCons,
		c.MinCons,
	)
	return connStr
}
