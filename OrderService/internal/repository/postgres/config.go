package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string `envconfig:"DATABASE_HOST" default:"localhost"`
	Port     int    `envconfig:"DATABASE_PORT" default:"5432"`
	Username string `envconfig:"DATABASE_USERNAME" default:"postgres"`
	Password string `envconfig:"DATABASE_PASSWORD" default:"postgres"`
	Name     string `envconfig:"DATABASE_NAME" default:"postgres"`
	SSLMode  string `envconfig:"DATABASE_SSLMODE" default:"disable"`
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

	//err = db.AutoMigrate()
	return db, nil
}
