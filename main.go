package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-sphero"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	spheroPort := os.Getenv("SPHERO_PORT")
	if spheroPort == "" {
		spheroPort = "/dev/tty.Sphero-ORY-AMP-SPP"
	}

	master := gobot.GobotMaster()
	gobot.Api(master)

	robots := map[string][]string{
		"gundam": []string{"sphero", spheroPort},
	}

	for name, spheros := range robots {
		spheroAdaptor := new(gobotSphero.SpheroAdaptor)
		spheroAdaptor.Name = spheros[0]
		spheroAdaptor.Port = spheros[1]

		sphero := gobotSphero.NewSphero(spheroAdaptor)
		sphero.Name = spheros[0]

		connections := []interface{}{
			spheroAdaptor,
		}
		devices := []interface{}{
			sphero,
		}

		robot := gobot.Robot{
			Name:        name,
			Connections: connections,
			Devices:     devices,
		}

		master.Robots = append(master.Robots, robot)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		for _, robot := range master.Robots {
			fmt.Printf("Stopping devise %s...\n", robot.Name)
			gobot.Call(robot.GetDevice(robot.Name).Driver, "StopC")
		}
		os.Exit(1)
	}()

	master.Start()
}
