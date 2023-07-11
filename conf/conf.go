package conf

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var fileCfg *Config

// Config struct for webapp config
type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Redis struct {
		Addr   string `yaml:"addr"`
		Passwd string `yaml:"passwd"`
	} `yaml:"redis"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) *Config {
	config := &Config{}
	file, err := os.Open(configPath)
	if err != nil {
		panic(fmt.Sprintf("NewConfig err:%s", err))
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		panic(fmt.Sprintf("NewConfig err:%s", err))
	}
	return config
}
func GetConfig() *Config {
	if fileCfg == nil {
		fileCfg = NewConfig("./conf/conf.yaml")
	}
	return fileCfg
}
