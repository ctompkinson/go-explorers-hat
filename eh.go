package explorers

import (
	"gobot.io/x/gobot/platforms/raspi"
	"gobot.io/x/gobot"
	"github.com/sirupsen/logrus"
)

type ExplorersHat struct {
	Log       *logrus.Logger
	Pi        *raspi.Adaptor
	Leds      []led
	Motors    []motor
	Inputs    []input
	Outputs   []output
	Analgoues []analogue
	Work      func(*ExplorersHat)
}

func Setup(work func(*ExplorersHat), logLevel logrus.Level) ExplorersHat {
	pi := raspi.NewAdaptor()
	eh := ExplorersHat{
		Log:  logrus.StandardLogger(),
		Pi:   pi,
		Work: work,
	}
	eh.Log.SetLevel(logLevel)
	eh.SetupDevices()
	return eh
}

func (eh *ExplorersHat) SetupDevices() {
	eh.Log.Debugln("Setting up explorers hat devices")

	eh.Leds = append(eh.Leds, eh.NewLed(PinLed1))
	eh.Leds = append(eh.Leds, eh.NewLed(PinLed2))
	eh.Leds = append(eh.Leds, eh.NewLed(PinLed3))
	eh.Leds = append(eh.Leds, eh.NewLed(PinLed4))

	eh.Motors = append(eh.Motors, eh.NewMotor(PinM1f, PinM1b))
	eh.Motors = append(eh.Motors, eh.NewMotor(PinM2f, PinM2b))

	eh.Inputs = append(eh.Inputs, eh.NewInput(PinIn1))
	eh.Inputs = append(eh.Inputs, eh.NewInput(PinIn2))
	eh.Inputs = append(eh.Inputs, eh.NewInput(PinIn3))
	eh.Inputs = append(eh.Inputs, eh.NewInput(PinIn4))

	eh.Outputs = append(eh.Outputs, eh.NewOutput(PinOut1))
	eh.Outputs = append(eh.Outputs, eh.NewOutput(PinOut2))
	eh.Outputs = append(eh.Outputs, eh.NewOutput(PinOut3))
	eh.Outputs = append(eh.Outputs, eh.NewOutput(PinOut4))
}

func (eh *ExplorersHat) Run() {

	DoWork := func() {
		eh.Work(eh)
	}

	eh.Log.Debugln("creating gobot robot")
	robot := gobot.NewRobot("ExplorersHat",
		[]gobot.Connection{eh.Pi},
		[]gobot.Device{
			eh.Leds[0].Device,
			eh.Leds[1].Device,
			eh.Leds[2].Device,
			eh.Leds[3].Device,

			eh.Inputs[0].Device,
			eh.Inputs[1].Device,
			eh.Inputs[2].Device,
			eh.Inputs[3].Device,

			eh.Outputs[0].device,
			eh.Outputs[1].device,
			eh.Outputs[2].device,
			eh.Outputs[3].device,
		},
		DoWork,
	)

	eh.Log.Debugln("stating gobot robot")
	robot.Start()
}
