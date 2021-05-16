package main

import (
	"fmt"
	"os"
)

type DataBaseConfig struct {
	Host     string
	User     string
	Password string
	Database string
}

func (c *DataBaseConfig)BuildDataSourcename() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s",c.User,c.Password,c.Host,c.Database)
}

func BuildDatabaseConfigFromEnvironment() *DataBaseConfig {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")

	return &DataBaseConfig{
		Host: dbHost,
		User: dbUser,
		Password: dbPassword,
		Database: dbDatabase,
	}
}