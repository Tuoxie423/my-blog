package core

import (
	"log"
	"server/config"
	"server/utils"

	"gopkg.in/yaml.v3"
)

func InitConf() *config.Config {
	c := &config.Config{}

	// 加载配置文件
	yamlConf, err := utils.LoadYaml()
	if err != nil {
		log.Fatalf("加载配置文件失败：%v", err)
	}
	if err := yaml.Unmarshal(yamlConf, c); err != nil {
		log.Fatalf("解析配置文件失败：%v", err)
	}
	return c

}
