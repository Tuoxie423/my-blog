package utils

import (
	"io/fs"
	"os"
	"server/global"

	"gopkg.in/yaml.v3"
)

const (
	path = "config.yaml"
)

func LoadYaml() ([]byte, error) {
	return os.ReadFile(path)
}

func SaveYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	return os.WriteFile(path, byteData, fs.ModePerm)
}
