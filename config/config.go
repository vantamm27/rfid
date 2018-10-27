package config

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"rfid/common/utility"

	"gopkg.in/natefinch/lumberjack.v2"
)

type CONFIG struct {
	AppName string     `json:"app"`
	Mqtt    MQTTCONFIG `json:"mqtt"`
	HID     HIDCONFIG  `json:"hid"`
	Log     LOG        `json:"log"`
}

type MQTTCONFIG struct {
	Url      string `json:"url"`
	Topic    string `json:"topic"`
	Ping     string `json:"ping"`
	Qos      byte   `json:"qos"`
	Retained bool   `json:"retained"`
}

type HIDCONFIG struct {
	DataLen    int    `json:"datalen"`
	DeviceType string `json:"deviceType"`
}

type LOG struct {
	Path   string `json:"path"`
	Size   int    `json:"size"`
	Age    int    `json:"age"`
	Backup int    `json:"backup"`
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

	configLog()
}

func configLog() {

	logDir := filepath.Dir(Config.Log.Path)
	log.Println(logDir)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0777)
	}
	var out io.Writer = &lumberjack.Logger{
		Filename:   Config.Log.Path,
		MaxSize:    5,
		MaxBackups: 10,
		MaxAge:     10,
		Compress:   false,
		LocalTime:  true,
	}

	mw := io.MultiWriter(out, os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	log.SetOutput(mw)

}
