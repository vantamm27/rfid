// main.go
package main

import (
	"log"
	"os"
	"os/signal"

	"rfid/common"
	"rfid/config"
	"rfid/myhid"
	"syscall"
)

const (
	ConfigFile string = "./conf/config.json"
)

func main() {
	config.Init(ConfigFile)
	common.InitMap()
	myhid.Init()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	select {
	case a := <-c:
		log.Printf("=== Recved sign %d\n", a)
		break
	}
	myhid.Close()

}
