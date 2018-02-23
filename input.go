package explorers

import (
	"gobot.io/x/gobot/drivers/gpio"
	"strconv"
)

type input struct {
	Device *gpio.DirectPinDriver
}

func (eh *ExplorersHat) NewInput(pin int) input {
	return input{gpio.NewDirectPinDriver(eh.Pi, strconv.Itoa(pin))}
}
