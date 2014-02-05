package main

import (
	"net/http"
	"strconv"
)

func NewApi(sphero Sphero) Api {
	return nil
}

type Api interface {
	Handler() http.Handler
}

func setRGB(s Sphero, r, g, b string) {
	rr, _ := strconv.Atoi(r)
	gg, _ := strconv.Atoi(g)
	bb, _ := strconv.Atoi(b)
	s.SetRGB((uint8)(rr), (uint8)(gg), (uint8)(bb))
}
