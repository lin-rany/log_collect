package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {
	conf := viper.New()
	pwd, _ := os.Getwd()
	fmt.Println(pwd)

	conf.AddConfigPath(pwd + "/conf")
	conf.SetConfigType("yaml")
	conf.SetConfigName("logagent")
	if err := conf.ReadInConfig(); err != nil {
		panic("ERR!")
	}

	fmt.Println(conf.Get("server"))
}
