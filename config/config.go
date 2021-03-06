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
		Port   int    `yaml:"port,omitempty"`
		Views  string `yaml:"views"`
		Assets string `yaml:"assets"`
	}
	Sql struct {
		MaxPageSize int `yaml:"maxPageSize"`
	}
	Mysql struct {
		User         string `yaml:"user"`
		Pwd          string `yaml:"pwd"`
		Port         int    `yaml:"port"`
		Host         string `yaml:"host"`
		Db           string `yaml:"db"`
		MaxIdleConns int    `yaml:"maxIdleConns"`
		MaxOpenConns int    `yaml:"maxOpenConns"`
	}
	Redis struct {
		Host string `yaml:"host"`
		Pwd  string `yaml:"pwd,omitempty"`
		Port int    `yaml:"port"`
		Db   int    `yaml:"db"`
	}
	Wxs struct {
		AppId     string `yaml:"appId"`
		AppSecret string `yaml:"appSecret"`
	}
	Wxm struct {
		AppId     string `yaml:"appId"`
		AppSecret string `yaml:"appSecret"`
	}
	WxOpen struct {
		AppId          string `yaml:"appId"`
		AppSecret      string `yaml:"appSecret"`
		ComponentPhone string `yaml:"componentPhone"`
	}
	WxMap struct {
		Key     string `yaml:"key"`
		Referer string `yaml:"referer"`
	}
	Oss struct {
		Domain          string `yaml:"domain"`
		Bucket          string `yaml:"bucket"`
		RegionId        string `yaml:"regionId"`
		AccessKeyId     string `yaml:"accessKeyId"`
		AccessKeySecret string `yaml:"accessKeySecret"`
		RoleArn         string `yaml:"roleArn"`

		DirectAccessKeyId     string `yaml:"directAccessKeyId"`
		DirectAccessKeySecret string `yaml:"directAccessKeySecret"`
	}
}

var configFile string

var conf Config

func init() {
	if strings.Contains(os.Args[0], "__Test") {
		flag.StringVar(&configFile, "f", "./../conf.yaml", "配置文件")
	} else {
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
	rlog.InfoF("%+v", conf)
}

func Cfg() Config {
	return conf
}
