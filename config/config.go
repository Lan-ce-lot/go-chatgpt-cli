package config

import (
	"go-chatgpt-cli/logger"
	"time"

	localconf "github.com/takecy/go-localconfig/conf"
)

// conf keep openAI API key
var conf *localconf.Conf
var Path string

// KeyConfig is config
type KeyConfig struct {
	Key  string
	Date time.Time
}

func Save(c interface{}) (err error) {
	err = conf.Save(c)
	return
}

func Load(c interface{}) (err error) {
	err = conf.Load(&c)
	return
}

func Cleanup() (err error) {
	err = conf.Cleanup()
	return
}

func init() {
	appName := "go-chatgpt-cli"
	fileName := "key.json"
	conf = localconf.NewConfig(appName, fileName)
	Path = conf.Path()
	logger.LanLogger.Printf("config.path: %s", conf.Path())
}
