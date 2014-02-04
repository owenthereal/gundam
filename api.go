package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/codegangsta/martini"
)

func NewApi(sphero Sphero) Api {
	return &ApiMartini{sphero}
}

type Api interface {
	Handler() http.Handler
}

type ApiMartini struct {
	S Sphero
}

func (a *ApiMartini) Start() {
	m := a.Handler()
	m.(*martini.ClassicMartini).Run()
}

func (a *ApiMartini) Handler() http.Handler {
	m := martini.Classic()
	m.Post("/rgb/:rgb", func(params martini.Params) (int, string) {
		err := setGRB(a.S, params["rgb"])
		if err != nil {
			return 400, fmt.Sprintf("%s", err)
		}

		return 201, "ok"
	})

	return m
}

func setGRB(sphero Sphero, rgb string) error {
	rgbRegexp := regexp.MustCompile(`^(\d+),(\d+),(\d+)$`)
	if !rgbRegexp.MatchString(rgb) {
		return fmt.Errorf("Invalid format of rgb")
	}

	r, _ := strconv.Atoi(rgbRegexp.FindStringSubmatch(rgb)[1])
	g, _ := strconv.Atoi(rgbRegexp.FindStringSubmatch(rgb)[2])
	b, _ := strconv.Atoi(rgbRegexp.FindStringSubmatch(rgb)[3])

	sphero.SetRGB((uint8)(r), (uint8)(g), (uint8)(b))

	return nil
}
