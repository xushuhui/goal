package config

import (
	"io/ioutil"
	"log"
	yaml "gopkg.in/yaml.v2"
)


type Conf struct {
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}

func ReadConfig() *Conf  {
	conf := new(Conf)
    yamlFile, err := ioutil.ReadFile("config/app.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, conf)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
	}
	return conf
}