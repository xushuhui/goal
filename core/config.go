package core

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Env struct {
	Env string `yaml:"env"`
}

type Conf struct {
	Mysql Mysql
	Redis Redis
}

type Mysql struct {
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}
type Redis struct {
	Host string `yaml:"host"`
}

var conf *Conf

const ENV = "dev"

func init() {
	yamlFile := readConf(ENV)
	if err := yaml.Unmarshal(yamlFile, &conf); err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
func readConf(fileName string) []byte {
	path := "./config/" + fileName + ".yaml"
	confPath, _ := filepath.Abs(path)
	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	return yamlFile
}
func ReadMysqlConf() Mysql {
	return conf.Mysql
}
func ReadRedisConf() Redis {
	return conf.Redis
}
