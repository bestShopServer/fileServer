package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type app struct {
	Name string
}

var AppSetting = &app{}

//本服务配置
type server struct {
	RunMode      string
	HttpAddr     string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	TokenExpire  int
	FilePath     string
	FileURL      string
}

var ServerSetting = &server{}

//腾讯云对象存储
type tencent_cos struct {
	AppID     string
	SecretID  string
	SecretKey string
	Region    string
}

var TencentSetting = &tencent_cos{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("tencent_cos", TencentSetting)

	//AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
