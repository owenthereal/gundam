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
	api        Api
	server     *httptest.Server
)

func setupMartini() {
	fakeSphero = &FakeSphero{}
	api = &ApiMartini{fakeSphero}
	server = httptest.NewServer(api.Handler())
}

func setupPlain() {
	fakeSphero = &FakeSphero{}
	api = &ApiPlain{fakeSphero}
	server = httptest.NewServer(api.Handler())
}

func tearDown() {
	server.Close()
}

func TestApiPlain_InvalidRgb(t *testing.T) {
	setupPlain()
	defer tearDown()

	url := fmt.Sprintf("%s/rgb/invalid", server.URL)
	req, _ := http.NewRequest("PUT", url, nil)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Fatal("err should be nil")
	}

	if resp.StatusCode != 404 {
		t.Fatalf("status code should be 404 but it's %s", resp.StatusCode)
	}
}

func TestApiPlain_ValidRgb(t *testing.T) {
	setupPlain()
	defer tearDown()

	url := fmt.Sprintf("%s/rgb/255,255,255", server.URL)
	req, _ := http.NewRequest("PUT", url, nil)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Fatal("err should be nil")
	}

	if resp.StatusCode != 201 {
		t.Fatalf("status code should be 404 but it's %s", resp.StatusCode)
	}

	if fakeSphero.R != (uint8)(255) {
		t.Fatalf("R should be 255 but it's %s", fakeSphero.R)
	}

	if fakeSphero.G != (uint8)(255) {
		t.Fatalf("G should be 255 but it's %s", fakeSphero.G)
	}

	if fakeSphero.B != (uint8)(255) {
		t.Fatalf("B should be 255 but it's %s", fakeSphero.B)
	}
}

func TestApiMartini_InvalidRgb(t *testing.T) {
	setupMartini()
	defer tearDown()

	url := fmt.Sprintf("%s/rgb/invalid", server.URL)
	req, _ := http.NewRequest("PUT", url, nil)
	resp, err := http.DefaultClient.Do(req)

	assert.Equal(t, nil, err)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestApiMartini_ValidRgb(t *testing.T) {
	setupMartini()
	defer tearDown()

	url := fmt.Sprintf("%s/rgb/255,255,255", server.URL)
	req, _ := http.NewRequest("PUT", url, nil)
	resp, err := http.DefaultClient.Do(req)

	assert.Equal(t, nil, err)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, (uint8)(255), fakeSphero.R)
	assert.Equal(t, (uint8)(255), fakeSphero.G)
	assert.Equal(t, (uint8)(255), fakeSphero.B)
}
