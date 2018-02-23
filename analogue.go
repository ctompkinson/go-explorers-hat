package explorers

import (
	"gobot.io/x/gobot/platforms/raspi"
	"strconv"
)

type analogue struct {
	Device *raspi.PWMPin
}

func (eh *ExplorersHat) NewAnalogue(pin int) analogue {
	return analogue{raspi.NewPWMPin(strconv.Itoa(pin))}
}
