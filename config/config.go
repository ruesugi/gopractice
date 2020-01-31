package config

import (
	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port   int
	Driver string
	Name   string
}

var Config ConfigList

func init() {
	c, err := ini.Load("config.ini")
	if err != nil {
		return
	}

	Config = ConfigList{
		Port:   c.Section("web").Key("port").MustInt(),
		Driver: c.Section("db").Key("driver").String(),
		Name:   c.Section("db").Key("name").MustString("stockdata.sql"),
	}
}
