package config

import (
	"fmt"

	"github.com/eduardogpg/gonv"
)

//DatabaseConfig struct
type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
	Debug    bool
}

var database *DatabaseConfig

func init() {
	database = &DatabaseConfig{}
	database.Username = gonv.GetStringEnv("USERNAME", "root")
	database.Password = gonv.GetStringEnv("PASSWORD", "")
	database.Host = gonv.GetStringEnv("HOST", "localhost")
	database.Port = gonv.GetIntEnv("PORT", 3306)
	database.Database = gonv.GetStringEnv("DATABASE", "participacion_ciudadana")
	database.Debug = gonv.GetBoolEnv("DEBUG", true)
}

//GetDebug method
func GetDebug() bool {
	return database.Debug
}

//GetURLDatabase method
func GetURLDatabase() string {
	return database.url()
}

func (d *DatabaseConfig) url() string {
	//  username:password@tcp(localhost:3066)/database?charset=utf8
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", d.Username, d.Password, d.Host, d.Port, d.Database)
}
