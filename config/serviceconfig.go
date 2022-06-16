package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ServiceConfig struct {
	ConsulIp    string         `yaml:"consulip"`
	NodesIP     []string       `yaml:"nodesip"`
	Sms         []ServiceModel `yaml:"services"`
	reliability bool           `yaml:"reliability"`
}

type ServiceModel struct {
	ConsulIp string `yaml:"consulip"`
	Name     string `yaml:"name"`
	Image    string `yaml:"image"`
	Cmd      string `yaml:"cmd"`
	Weight   int    `yaml:"weight"`
}

var sc ServiceConfig

func ConfigInit() (map[string]ServiceModel, string, map[string]bool, bool) {
	fileName := "config/config.yml"
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("fail to read file: ", err)
	}
	err = yaml.Unmarshal(file, &sc)
	if err != nil {
		log.Fatal("fail to yaml unmarshal: ", err)
	}

	var res map[string]ServiceModel = make(map[string]ServiceModel)
	for _, sm := range sc.Sms {
		res[sm.Name] = sm
	}

	var nodeStatus map[string]bool = make(map[string]bool)
	for _, ip := range sc.NodesIP {
		nodeStatus[ip] = true
	}

	return res, sc.ConsulIp, nodeStatus, sc.reliability
}
