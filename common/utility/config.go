package utility

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

func ReadConfiguration(fileName string, config interface{}) error {
	buff, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buff, config)
	if err != nil {
		return err
	}
	return nil
}

func GetCurrentDateTime() string {
	return time.Now().Format("20060102150405")
}
