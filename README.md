Gundam
======

Gundam is the most destructive robot! For now, it conquer the world with [Sphero](http://www.gosphero.com/).

![](https://dl.dropboxusercontent.com/u/1079131/00_gundam.jpg)

```plain
$ SPHERO_PORT=/dev/tty.Sphero-ORY-AMP-SPP go run main.go

$ curl http://localhost:3000/robots
[{"Connections":[{"Name":"sphero","Port":"/dev/tty.Sphero-ORY-AMP-SPP","Connected":false,"Params":null}],"Devices":[{"Interval":"0.1s","Pin":"","Name":"sphero","Params":null,"Commands":["SetRGBC","RollC","StopC"],"SpheroAdaptor":{"Name":"sphero","Port":"/dev/tty.Sphero-ORY-AMP-SPP","Connected":false,"Params":null}}],"Name":"gundam","Commands":null}]

$ curl http://localhost:3000/robots/gundam
{"Connections":[{"Name":"sphero","Port":"/dev/tty.Sphero-ORY-AMP-SPP","Connected":false,"Params":null}],"Devices":[{"Interval":"0.1s","Pin":"","Name":"sphero","Params":null,"Commands":["SetRGBC","RollC","StopC"],"SpheroAdaptor":{"Name":"sphero","Port":"/dev/tty.Sphero-ORY-AMP-SPP","Connected":false,"Params":null}}],"Name":"gundam","Commands":null}

$ curl http://localhost:3000/robots/gundam/devices
[{"Name":"sphero","Interval":"","Driver":{"Interval":"0.1s","Pin":"","Name":"sphero","Params":null,"Commands":["SetRGBC","RollC","StopC"],"SpheroAdaptor":{"Name":"sphero","Port":"/dev/tty.Sphero-ORY-AMP-SPP","Connected":false,"Params":null}},"Params":null}]

$ curl http://localhost:3000/robots/gundam/devices/sphero
{"Name":"sphero","Interval":"","Driver":{"Interval":"0.1s","Pin":"","Name":"sphero","Params":null,"Commands":["SetRGBC","RollC","StopC"],"SpheroAdaptor":{"Name":"sphero","Port":"/dev/tty.Sphero-ORY-AMP-SPP","Connected":false,"Params":null}},"Params":null}

$ curl http://localhost:3000/robots/gundam/devices/sphero/commands
["SetRGBC","RollC","StopC"]

$ curl -d '{"r": 0, "g": 255, "b": 0}' http://localhost:3000/robots/gundam/devices/sphero/commands/SetRGBC
{"results":[]}
```
