package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-sphero"
)

func NewSphero(name, port string) Sphero {
	return &sphero{Name: "Gundam", Port: port}
}

type Sphero interface {
	Start()
	Stop()
	SetRGB(r, g, b uint8)
}

type sphero struct {
	Name       string
	Port       string
	device     *gobotSphero.SpheroDriver
	connection *gobotSphero.SpheroAdaptor
}

func (s *sphero) Start() {
	s.connection = new(gobotSphero.SpheroAdaptor)
	s.connection.Name = "sphero"
	s.connection.Port = s.Port

	s.device = gobotSphero.NewSphero(s.connection)
	s.device.Name = s.Name

	robot := gobot.Robot{
		Connections: []gobot.Connection{s.connection},
		Devices:     []gobot.Device{s.device},
	}
	robot.Start()
}

func (s *sphero) Stop() {
	s.connection.Disconnect()
}

func (s *sphero) SetRGB(r, g, b uint8) {
	s.device.SetRGB(r, g, b)
}
