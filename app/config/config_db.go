package config

import (
	"fmt"
	"time"
)

type DBConfig struct {
	Host            string        `yaml:"host"`
	User            string        `yaml:"user"`
	Port            int           `yaml:"port"`
	Password        string        `yaml:"password"` // 这里会包含环境变量引用
	DBName          string        `yaml:"dbname"`
	SSLMode         string        `yaml:"sslmode"`
	MaxOpenConn     int           `yaml:"max_open_conn"`
	MaxIdleConn     int           `yaml:"max_idle_conn"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

func (d *DBConfig) ConnUrl() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.DBName, d.SSLMode)
}
