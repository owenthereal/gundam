package main

import (
	"regexp"
	"strconv"

	"github.com/codegangsta/martini"
)

type Api struct {
	S Sphero
}

func (a *Api) Start() {
	m := a.mux()
	m.Run()
}

func (a *Api) mux() *martini.ClassicMartini {
	m := martini.Classic()
	m.Post("/rgb/:rgb", func(params martini.Params) (int, string) {
		rgb := params["rgb"]
		rgbRegexp := regexp.MustCompile(`^(\d+),(\d+),(\d+)$`)
		if !rgbRegexp.MatchString(rgb) {
			return 400, "Invalid format of rgb"
		}

		r, _ := strconv.Atoi(rgbRegexp.FindStringSubmatch(rgb)[1])
		g, _ := strconv.Atoi(rgbRegexp.FindStringSubmatch(rgb)[2])
		b, _ := strconv.Atoi(rgbRegexp.FindStringSubmatch(rgb)[3])

		a.S.SetRGB((uint8)(r), (uint8)(g), (uint8)(b))

		return 201, "ok"
	})

	return m
}
