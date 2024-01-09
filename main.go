package main

import (
	"log"

	"github.com/song940/smartctl-go/smartctl"
)

func main() {
	dev := smartctl.Open("/dev/disk1s1")
	info, err := dev.Read()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(info.ModelName)
	log.Println(info.SerialNumber)
	log.Println(info.SmartStatus.Passed)
	log.Println(info.Temperature.Current)
	log.Println(info.PowerCycleCount)
}
