package main

import "os"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "/dev/tty.Sphero-ORY-AMP-SPP"
	}

	s := &sphero{Name: "Gundam", Port: port}
	defer s.Stop()

	api := &Api{s}
	go api.Start()

	s.Start()
}
