package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// Config 结构体定义
type Config struct {
	DB     DBConfig     `yaml:"db"`
	Server ServerConfig `yaml:"server"`
}

func Init() *Config {
	configPath := flag.String("c", "config.yaml", "配置文件路径") // 默认值为 config.yaml
	flag.Parse()

	// 1. 读取和解析配置文件
	config, err := loadConfig(*configPath)
	if err != nil {
		fmt.Printf("加载配置失败: %v\n", err)
		os.Exit(1)
	}
	return config
}

// loadConfig 加载并解析配置文件
func loadConfig(filename string) (*Config, error) {
	// 读取原始文件内容
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 处理环境变量替换
	expandedData := []byte(os.ExpandEnv(string(data)))

	// 使用Viper解析配置
	viper.SetConfigType("yml")
	if err := viper.ReadConfig(strings.NewReader(string(expandedData))); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}

	// 解析到结构体
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("反序列化配置失败: %w", err)
	}

	return &config, nil
}
