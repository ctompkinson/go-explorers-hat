package explorers

import (
	"gobot.io/x/gobot/drivers/gpio"
	"strconv"
)

type output struct {
	device *gpio.DirectPinDriver
}

func (eh *ExplorersHat) NewOutput(pin int) output {
	return output{gpio.NewDirectPinDriver(eh.Pi, strconv.Itoa(pin))}
}
