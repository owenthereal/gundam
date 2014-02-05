package main

import (
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

type ApiPlain struct {
	S Sphero
}

func (a *ApiPlain) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	rgbRegexp := regexp.MustCompile(`/rgb/(\d+),(\d+),(\d+)`)
	path := req.URL.Path
	if req.Method == "PUT" && rgbRegexp.MatchString(path) {
		match := rgbRegexp.FindStringSubmatch(path)
		setRGB(a.S, match[1], match[2], match[3])

		resp.WriteHeader(201)
	} else {
		http.NotFound(resp, req)
	}
}

func (a *ApiPlain) Handler() http.Handler {
	return a
}

type ApiMartini struct {
	S Sphero
}

func (a *ApiMartini) Handler() http.Handler {
	m := martini.Classic()
	m.Put("/rgb/:rgb", func(params martini.Params) (int, string) {
		rgb := params["rgb"]

		rgbRegexp := regexp.MustCompile(`^(\d+),(\d+),(\d+)$`)
		if !rgbRegexp.MatchString(params["rgb"]) {
			return 400, "Invalid format of rgb"
		}

		match := rgbRegexp.FindStringSubmatch(rgb)
		setRGB(a.S, match[1], match[2], match[3])

		return 201, "ok"
	})

	return m
}

func setRGB(s Sphero, r, g, b string) {
	rr, _ := strconv.Atoi(r)
	gg, _ := strconv.Atoi(g)
	bb, _ := strconv.Atoi(b)
	s.SetRGB((uint8)(rr), (uint8)(gg), (uint8)(bb))
}
