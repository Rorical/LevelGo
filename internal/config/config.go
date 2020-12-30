package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type LevelDBSetting struct {
	File    string
	RpcPort uint
}

func Read() *LevelDBSetting {
	var leveldbConf LevelDBSetting
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("json")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	if err := v.Unmarshal(&leveldbConf); err != nil {
		fmt.Println(err)
	}
	return &leveldbConf
}
