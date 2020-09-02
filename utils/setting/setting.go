// File:    setting
// Version: 1.0.0
// Creator: JoeLang
// Date:    2020/8/30 21:38
// DESC:    load setting message

package setting

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Host     string
	Port     string
	UserName string
	Password string
	DataBase string
)

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("Load Initialization Config Failed.")
	}
	loadServer(cfg)
	loadMysql(cfg)
}

func loadServer(cfg *ini.File) {
	AppMode = cfg.Section("server").Key("AppMode").MustString("debug")
	HttpPort = cfg.Section("server").Key("HttpPort").MustString(":8888")
}

func loadMysql(cfg *ini.File) {
	Host = cfg.Section("mysql").Key("Host").MustString("")
	Port = cfg.Section("mysql").Key("Port").MustString("")
	UserName = cfg.Section("mysql").Key("UserName").MustString("")
	Password = cfg.Section("mysql").Key("Password").MustString("")
	DataBase = cfg.Section("mysql").Key("DataBase").MustString("")
}
