package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-sphero"
)

type Sphero struct {
	Name   string
	Port   string
	device *gobotSphero.SpheroDriver
}

func (s *Sphero) Start() {
	spheroAdaptor := new(gobotSphero.SpheroAdaptor)
	spheroAdaptor.Name = "sphero"
	spheroAdaptor.Port = s.Port

	s.device = gobotSphero.NewSphero(spheroAdaptor)
	s.device.Name = s.Name

	robot := gobot.Robot{
		Connections: []gobot.Connection{spheroAdaptor},
		Devices:     []gobot.Device{s.device},
	}
	robot.Start()
}
