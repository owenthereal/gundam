package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
)

type FakeSphero struct {
	R, G, B uint8
}

func (s *FakeSphero) Start() {}
func (s *FakeSphero) Stop()  {}
func (s *FakeSphero) SetRGB(r, g, b uint8) {
	s.R, s.G, s.B = r, g, b
}

var (
	fakeSphero *FakeSphero
	api        *Api
	server     *httptest.Server
)

func setup() {
	fakeSphero = &FakeSphero{}
	api = &Api{fakeSphero}
	server = httptest.NewServer(api.mux())
}

func tearDown() {
	server.Close()
}

func TestApi_InvalidRgb(t *testing.T) {
	setup()
	defer tearDown()

	url := fmt.Sprintf("%s/rgb/invalid", server.URL)
	resp, err := http.Post(url, "", nil)

	assert.Equal(t, nil, err)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestApi_ValidRgb(t *testing.T) {
	setup()
	defer tearDown()

	url := fmt.Sprintf("%s/rgb/255,255,255", server.URL)
	resp, err := http.Post(url, "", nil)

	assert.Equal(t, nil, err)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, (uint8)(255), fakeSphero.R)
	assert.Equal(t, (uint8)(255), fakeSphero.G)
	assert.Equal(t, (uint8)(255), fakeSphero.B)
}
