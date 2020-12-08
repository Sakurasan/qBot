package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

// GlobalConfig 默认全局配置
var GlobalConfig *Config

// Init 使用 ./application.yaml 初始化全局配置
func Init() {
	err := checkFile()
	if err != nil {
		logrus.WithField("config", "checkFile").WithError(err).Panicln("no config file")
	}
	GlobalConfig = &Config{
		viper.New(),
	}
	GlobalConfig.SetConfigName("qBot")
	GlobalConfig.SetConfigType("yaml")
	GlobalConfig.AddConfigPath(".")
	GlobalConfig.AddConfigPath("./config")

	err = GlobalConfig.ReadInConfig()
	if err != nil {
		logrus.WithField("config", "GlobalConfig").WithError(err).Panicf("unable to read global config")
	}
}

func checkFile() error {
	// _, err := os.Stat("./config")
	// if err != nil {
	// 	os.MkdirAll("config", os.ModeDir|0644)
	// }
	if !IsFileExist("./qBot.yaml") {
		f, _ := os.OpenFile("./qBot.yaml", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		defer f.Close()
		_, err := f.WriteString(template)
		if err != nil {
			return err
		}
	}
	return nil
}

//判断文件是否存在
func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

var template = `
account : 
pwd : 

groupID: 
  - 
`
