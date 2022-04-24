package main

import "fmt"

type LegacyDevice interface {
	Do(s string) string
}
type legacyDevice struct {}

func (l *legacyDevice) Do(s string) string {
	fmt.Printf("Legacy device %s", s)
	return "ok"
}

type ModernDevice interface {
	DoStored() string
}

type DeviceAdapter struct {
	OldDevice LegacyDevice
	Msg string
}

func (a *DeviceAdapter) DoStored() string {
	var newMsg string
	if a.OldDevice != nil {
		newMsg = a.OldDevice.Do("s")
	} else {
		newMsg = a.Msg
	}
	return newMsg
}