package postgres

import (
	"OrderService/internal/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string `env:"POSTGRES_HOST" default:"localhost"`
	Port     int    `env:"POSTGRES_PORT" default:"5432"`
	Username string `env:"POSTGRES_USERNAME" default:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" default:"postgres"`
	Name     string `env:"POSTGRES_NAME" default:"postgres"`
	SSLMode  string `env:"POSTGRES_SSLMODE" default:"disable"`
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		c.Host,
		c.Username,
		c.Password,
		c.Name,
		c.Port,
		c.SSLMode,
	)
}

func NewDBConfig(cfg Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&model.Order{},
		&model.Delivery{},
		&model.Payment{},
		&model.Item{}); err != nil {
		return nil, err
	}
	return db, nil
}
