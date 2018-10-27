package myhid

import (
	"log"
	"rfid/common"
	"rfid/common/utility"
	"rfid/config"
	"rfid/mqtt"
	"strings"
	"time"

	"github.com/karalabe/hid"
)

const (
	DEVICE_TYPE_READER = "reader"
)

var (
	deviceList []*hid.Device
	mqttClient mqtt.MqttClient
	bStop      bool = false
)

func Init() {
	//dvs := hid.Enumerate(65535, 53) // one devivce

	initHID()
	err := initMQTT()
	if err != nil {
		log.Fatalln("myhid.initMQTT "+"Connect to mqtt falure ", err.Error())
	}
	ping()
}

func initHID() {
	log.Println("initHID")
	dvs := hid.Enumerate(0, 0)
	for _, dv := range dvs {

		if strings.Contains(strings.ToLower(dv.Product), config.Config.HID.DeviceType) && dv.Interface == 0 {
			log.Printf("myhid.initMQTT "+"%+v", dv)
			log.Println("myhid.initMQTT >>>" + "open device")

			go readData(dv)
		}
	}
}

func Close() {
	bStop = true

	for _, dv := range deviceList {
		log.Printf("myhid.myhid"+"close device %+v", dv)
		err := dv.Close()
		if err != nil {
			log.Printf("myhid.myhid"+"close device error", err.Error())
		}

	}
	mqttClient.Close()
}

func readData(dvInfo hid.DeviceInfo) {
	log.Printf("myhid.myhid"+"Open device %+v", dvInfo)
	dv, err := dvInfo.Open()
	deviceList = append(deviceList, dv)
	if err != nil {
		log.Printf("myhid.myhid"+"Open device error %+v", dvInfo)
		log.Println("myhid.myhid " + err.Error())
		return
	}
	log.Println("myhid.myhid " + "Open device success")
	var buff []byte
	var rfid string = ""
	buff = make([]byte, config.Config.HID.DataLen)
	var n int = 0

	for {
		n, err = dv.Read(buff)
		if err != nil {
			if bStop {
				dv.Close()
				dv = nil
				log.Println("myhid.myhid " + "Close device")
				return
			}
			log.Println("myhid.myhid "+"read data from device error", err.Error())
			break
		} else if n > 0 {
			if buff[2] > 0 {
				rfid += common.GetCharacter(buff[2])
			}

			if len(rfid) == config.Config.HID.DataLen {
				log.Println("myhid.myhid" + rfid)
				process(rfid)
				rfid = ""
			}
			continue
		}
		if bStop {
			return
		}

	}
	log.Println("++++++++++++")
}

func initMQTT() error {
	var err error
	mqttClient, err = mqtt.NewClient(config.Config.Mqtt.Url, config.Config.AppName+"_"+utility.GetCurrentDateTime())

	if err != nil {
		log.Println("myhid.initMQTT"+"Init mqtt connection error", err.Error())
		return err
	}

	err = mqttClient.Connect()
	if err != nil {
		log.Println("myhid.initMQTT "+"connect to mqtt server error", err.Error())
		return err
	}

	return nil
}

func process(data string) {
	log.Println("myhid.process " + data)
	err := mqttClient.Publish(config.Config.Mqtt.Topic, config.Config.Mqtt.Qos, config.Config.Mqtt.Retained, []byte(data))
	if err != nil {
		log.Println("pub data to mqtt err", err.Error())
		return
	}
}

func ping() {
	tickChan := time.NewTicker(3 * time.Second).C
	for {
		select {
		case <-tickChan:
			log.Println("myhid.ping ")
			if mqttClient == nil {
				log.Fatalln("mqttClient is null")
			}
			err := mqttClient.Publish(config.Config.Mqtt.Ping, config.Config.Mqtt.Qos, config.Config.Mqtt.Retained, []byte(config.Config.AppName))
			if err != nil {
				log.Println("pub data to mqtt err", err.Error())

			}
			if len(deviceList) <= 0 {
				initHID()
			}
		}
	}
}
