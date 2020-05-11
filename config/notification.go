package config


import "github.com/spf13/viper"

type Message struct {
	Appkey   string
	Tplid     string
	Sdkappid     string
}

func InitMessage(cfg *viper.Viper) *Message {
	return &Message{
		Appkey:     cfg.GetString("appkey"),
		Tplid:   cfg.GetString("tplid"),
		Sdkappid:     cfg.GetString("sdkappid"),
	}
}

var MessageConfig = new(Message)