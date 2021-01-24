package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"mylogtransfer/es"
	"mylogtransfer/kafka"
	"mylogtransfer/model"
)

func main() {
	
	var cfg = new(model.Config)
	err := ini.MapTo(cfg, "./config/logtransfer.ini")
	if err != nil {
		fmt.Println("load config failed, err:", err)
		
		panic(err)
	}
	fmt.Println("load config success=", cfg)

	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.Topic)
	if err != nil {
		fmt.Println("connect to kafka failed, err=", err)
		panic(err)
	}
	
	err = es.Init(cfg.ESConf.Address, cfg.ESConf.Index, cfg.ESConf.GoNum, cfg.ESConf.MaxSize)
	if err != nil {
		fmt.Println("Init es failed, err=", err)
		panic(err)
	}
	fmt.Println("Inir es success")
	select {}

}
