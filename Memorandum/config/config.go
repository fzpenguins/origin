package config

import (
	"github.com/go-ini/ini"
	"log"
	"strings"
)

var (
	dbUser      string
	dbPassword  string
	dbTableName string
)

func Init() {
	IniFile, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Println("err = ", err)
		return
	}
	getValue(IniFile)
	dsn := strings.Join([]string{dbUser, ":", dbPassword, "@tcp(localhost:3306)/", dbTableName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	InitDB(dsn)
}

func getValue(IniFile *ini.File) {
	dbUser = IniFile.Section("database").Key("dbUser").String()
	dbPassword = IniFile.Section("database").Key("dbPassword").String()
	dbTableName = IniFile.Section("database").Key("dbTableName").String()
}
