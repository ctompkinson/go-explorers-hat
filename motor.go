package explorers

import (
	"gobot.io/x/gobot/platforms/raspi"
	"strconv"
)

type motor struct {
	forwardPin  *raspi.PWMPin
	backwardPin *raspi.PWMPin
}

func (eh *ExplorersHat) NewMotor(forwardPin int, backwardPin int) (motor) {
	return motor{
		raspi.NewPWMPin(strconv.Itoa(forwardPin)),
		raspi.NewPWMPin(strconv.Itoa(backwardPin)),
	}
}

func (m *motor) Forward() {
	m.forwardPin.SetDutyCycle(1)
	m.backwardPin.SetDutyCycle(0)
}

func (m *motor) Backward() {
	m.forwardPin.SetDutyCycle(0)
	m.backwardPin.SetDutyCycle(1)
}
