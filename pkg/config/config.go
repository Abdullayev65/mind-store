package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"net/url"
	"os"
	"path"
	"strings"
)

var config = new(Config)

type Config struct {
	BaseUrl   string    `yaml:"base_url"`
	Port      string    `yaml:"port"`
	DebugMode bool      `yaml:"debug_mode"`
	DB        *DB       `yaml:"db"`
	JwtToken  *JwtToken `yaml:"jwt_token"`
}

func (cnfg *Config) GetPort() string {
	if strings.HasPrefix(cnfg.Port, ":") {
		return cnfg.Port
	}

	return ":" + cnfg.Port
}

func GetConfig() Config {
	return *config
}
func GetDB() DB {
	return *config.DB
}
func GetJwtToken() JwtToken {
	return *config.JwtToken
}

func GetPort() string {
	port := config.Port
	if port == "" {
		return ":8080"
	}
	if port[0] == ':' {
		return port
	}
	return ":" + port
}

func GetFilesBaseUrl() *url.URL {
	url, _ := url.Parse(config.BaseUrl + "/api/v1/file/")

	return url
}
func GetFilesUrlWith(suffixUrl string) string {
	url := GetFilesBaseUrl()

	url.Path = path.Join(url.Path, suffixUrl)

	return url.String()
}

func unmarshalConfig() (*Config, error) {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("yamlFile.Get err   #%v ", err)
	}

	c := new(Config)
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, fmt.Errorf("yaml.Unmarshal: %v", err)
	}

	return c, nil
}

func init() {
	c, err := unmarshalConfig()
	if err != nil {
		panic(err)
	}

	*config = *c
}
