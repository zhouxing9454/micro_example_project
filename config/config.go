package config

import (
	"fmt"
	"sync"

	"gopkg.in/ini.v1"
)

var (
	wg                  sync.WaitGroup
	EtcdHost            string
	EtcdPort            string
	HttpHost            string
	HttpPort            string
	HelloServiceAddress string
)

func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Printf("load config failed,err:%v\n", err)
	}
	wg.Add(3)
	go LoadEtcd(file)
	go LoadGin(file)
	go LoadService(file)
	wg.Wait()
}

func LoadEtcd(file *ini.File) {
	EtcdHost = file.Section("Etcd").Key("EtcdHost").String()
	EtcdPort = file.Section("Etcd").Key("EtcdPort").String()
	wg.Done()
}

func LoadGin(file *ini.File) {
	HttpHost = file.Section("Gin").Key("HttpHost").String()
	HttpPort = file.Section("Gin").Key("HttpPort").String()
	wg.Done()
}

func LoadService(file *ini.File) {
	HelloServiceAddress = file.Section("Service").Key("HelloServiceAddress").String()
	wg.Done()
}
