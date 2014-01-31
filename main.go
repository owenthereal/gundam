package main

import "os"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "/dev/tty.Sphero-ORY-AMP-SPP"
	}

	sphero := &Sphero{Name: "Gundam", Port: port}
	sphero.Start()
}
