package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
	"wx-server/rlog"
)

type Config struct {
	Web struct {
		Port int `yaml:"port,omitempty"`
	}
	Mysql struct {
		User string `yaml:"user"`
		Pwd  string `yaml:"pwd"`
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
		Db   string `yaml:"db"`
		MaxIdleConns int `yaml:"maxIdleConns"`
		MaxOpenConns int `yaml:"maxOpenConns"`
	}
	Redis struct {
		Host string `yaml:"host"`
		Pwd  string `yaml:"pwd,omitempty"`
		Port int    `yaml:"port"`
		Db   int    `yaml:"db"`
	}
	Wxs struct{
		AppId string `yaml:"appId"`
		AppSecret string `yaml:"appSecret"`
	}
	WxMap struct{
		Key string `yaml:"key"`
	}
}

var configFile string

var conf Config

func init() {
	if strings.Contains(os.Args[0],"__Test"){
		flag.StringVar(&configFile, "f", "./../conf.yaml", "配置文件")
	}else{
		flag.StringVar(&configFile, "f", "./conf.yaml", "配置文件")
	}

	f, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(configFile + " 不存在")
	}
	err = yaml.Unmarshal(f, &conf)
	if err != nil {
		fmt.Println("解析错误")
	}
	rlog.Info(conf)
}

func Cfg() Config {
	return conf
}
