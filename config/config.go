package config

import (
	"log"
	"rfid/common/utility"
)

type CONFIG struct {
	AppName string     `json:"app"`
	Mqtt    MQTTCONFIG `json:"mqtt"`
	HID     HIDCONFIG  `json:"hid"`
}

type MQTTCONFIG struct {
	Url      string `json:"url"`
	Topic    string `json:"topic"`
	Qos      byte   `json:"qos"`
	Retained bool   `json:"retained"`
}

type HIDCONFIG struct {
	DataLen    int    `json:"datalen"`
	DeviceType string `json:"deviceType"`
}

var (
	Config CONFIG
)

func Init(path string) {
	err := utility.ReadConfiguration(path, &Config)
	if err != nil {
		log.Fatal("config.Init", err.Error())
		return
	}
	log.Printf("config.Init %+v", Config)
}
