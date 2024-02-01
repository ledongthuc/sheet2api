package configs

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	HostIP   string `yaml:"host_ip"`
	HostPort string `yaml:"host_port"`
	Files    []File `yaml:"files"`
}

type File struct {
	FilePath        string `yaml:"file_path"`
	URLReplacedName string `yaml:"api_replaced_name"`
}

func LoadConfigFile(filePath string) (*Config, error) {
	config := &Config{
		HostIP:   DefaultHostIP,
		HostPort: DefaultPort,
	}
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal([]byte(fileContent), config)
	if err != nil {
		return config, fmt.Errorf("unmarshal configs: %w", err)
	}

	return config, nil
}

func GetEnvConfig(envName, defaultValue string) string {
	if val := os.Getenv(envName); len(val) > 0 {
		return val
	}
	return defaultValue
}
