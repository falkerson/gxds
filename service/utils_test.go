package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/tonyespy/gxds"
)

func TestConsulClientReturnErrorOnTimeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Second * 11)
		}))
	defer ts.Close()

	url := strings.Split(ts.URL, ":")
	host := url[0] + ":" + url[1]
	port, err := strconv.Atoi(url[2])
	if err != nil {
		fmt.Println(err.Error())
	}

	config := &gxds.Config{}
	config.Registry.Host = host
	config.Registry.Port = port
	config.Registry.FailLimit = 1
	config.Registry.FailWaitTime = 0

	consul, err := getConsulClient(config)
	if consul != nil || err == nil {
		t.Error("Error should be raised")
	}

	if err.Error() != "Cannot get connection to Consul" {
		t.Error("Wrong error message")
	}
}

func TestConsulClientReturnErrorOnBadResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusServiceUnavailable)
		}))
	defer ts.Close()

	url := strings.Split(ts.URL, ":")
	host := url[0] + ":" + url[1]
	port, err := strconv.Atoi(url[2])
	if err != nil {
		fmt.Println(err.Error())
	}

	config := &gxds.Config{}
	config.Registry.Host = host
	config.Registry.Port = port
	config.Registry.FailLimit = 1
	config.Registry.FailWaitTime = 0

	consul, err := getConsulClient(config)
	if consul != nil || err == nil {
		t.Error("Error should be raised")
	}

	if err.Error() != "Bad response from Consul service" {
		t.Error("Wrong error message")
	}
}

func TestLoadConfig(t *testing.T) {}

func TestBuildAddr(t *testing.T) {
	addr := buildAddr("test_host", "1000")
	if addr != "http://test_host:1000" {
		t.Error("Wrong URL returned")
	}
}

func TestCompareStrings(t *testing.T) {
	strings1 := []string{"one", "two", "three"}
	strings2 := []string{"one", "two"}
	strings3 := []string{"one", "two", "THREE"}

	if !compareStrings(strings1, strings1) {
		t.Error("Equal slices fail check!")
	}

	if compareStrings(strings1, strings2) {
		t.Error("Different size slices are OK!")
	}

	if compareStrings(strings1, strings3) {
		t.Error("Slice with different strings are OK!")
	}
}

func TestCompareStrStrMap(t *testing.T) {
	map1 := map[string]string{
		"album":  "electric ladyland",
		"artist": "jimi hendrix",
		"guitar": "white strat",
	}

	map2 := map[string]string{
		"album":  "electric ladyland",
		"artist": "jimi hendrix",
	}

	map3 := map[string]string{
		"album":  "iv",
		"artist": "led zeppelin",
		"guitar": "les paul",
	}

	if !compareStrStrMap(map1, map1) {
		t.Error("Equal maps fail check")
	}

	if compareStrStrMap(map1, map2) {
		t.Error("Different size maps are OK!")
	}

	if compareStrStrMap(map1, map3) {
		t.Error("Maps with different content are OK!")
	}
}
