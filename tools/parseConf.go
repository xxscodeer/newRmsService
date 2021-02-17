package tools

import (
	"io/ioutil"
	"log"
	yaml "gopkg.in/yaml.v2"
)
type Config struct {
	App    AppConfig    `yaml:"app"`
	Mysql  MysqlConfig  `yaml:"mysql"`
}
type AppConfig struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Mode    string `yaml:"mode"`
}

type MysqlConfig struct {
	DbName    string `yaml:"dbName"`
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	DbUser    string `yaml:"dbUser"`
	DbPwd     string `yaml:"dbPwd"`
	Loc       string `yaml:"loc"`
	ParseTime string `yaml:"parseTime"`
	Charset   string `yaml:"charset"`
}
func ParseConfig(path string) (cfg *Config) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("open file err,", err)
	}
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		log.Fatalln("unmarshal fail err,", err)
	}
	return
}
