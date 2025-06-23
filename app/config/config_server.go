package config

import "fmt"

type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

func (c *ServerConfig) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
