package explorers

import (
	"gobot.io/x/gobot/drivers/gpio"
	"strconv"
)

type led struct {
	eh     *ExplorersHat
	Device *gpio.LedDriver
}

func (eh *ExplorersHat) NewLed(ledPin int) led {
	eh.Log.Debugln("Creating new LED")
	return led{eh,gpio.NewLedDriver(eh.Pi, strconv.Itoa(ledPin))}
}

func (l *led) Toggle() {
	l.eh.Log.Debugln("Toggling LED on pin " + l.Device.Pin())
	l.Device.Toggle()
}
