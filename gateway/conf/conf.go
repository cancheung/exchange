package conf

import (
	"flag"
	"github.com/Unknwon/goconfig"
)

var (
	confPath string
	Cfg      *goconfig.ConfigFile
	AppKey string


)

func NewConfig(path string) *goconfig.ConfigFile {
	ConfigFile, err := goconfig.LoadConfigFile(path)
	if err != nil {
		panic("load config err is " + err.Error())
		return nil
	}
	return ConfigFile
}

func init() {
	flag.StringVar(&confPath, "conf", "gateway.ini", "config path")

}

func Init() {
	Cfg = NewConfig(confPath)
	AppKey = Cfg.MustValue("app", "app_key", "pfdsapowmsapa")
}
