package main

/**
	Chainsaw, a web audit tool.
 */

import (
	"bufio"
	"chainsaw/baseline"
	"chainsaw/tools"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var argFile = flag.String("f", "", "path to file")

type Proxy struct {
}

func (p *Proxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Received request %s %s %s\n", req.Method, req.Host, req.RemoteAddr)
	transport :=  http.DefaultTransport
	outReq := new(http.Request)
	*outReq = *req
	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}
	for key, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(key, v)
		}
	}
	rw.WriteHeader(res.StatusCode)
	_, _ = io.Copy(rw, res.Body)
	res.Body.Close()
}

func main() {
	flag.Parse()
	if len(os.Args) <= 1 {
		fmt.Println("[*] Use -help to get help.")
		os.Exit(0)
	}
	if *argFile != "" {
		file, err := os.Open(*argFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			core(scanner.Text())
		}
		os.Exit(0)
	}
	u := os.Args[1]
	core(u)
	http.Handle("/", &Proxy{})
	_ = http.ListenAndServe("0.0.0.0:1234", nil)
}

func core(urlText string) {
	fmt.Println("[+] Working on "+ urlText +"...")
	Url, _ := url.Parse(urlText)
	checkUrl(Url)
	if isAlive(Url) {
		baseline.Start(Url)
	} else {
		log.Println("[*] " + Url.String() + " not alive!")
	}
	fmt.Println("[+] Done.")
}

func isAlive(Url *url.URL) bool {
	resp := tools.DoRequest(Url, tools.ReqParam{Timeout: 10 * time.Second})
	defer resp.Body.Close()
	return true
}

func checkUrl(url *url.URL) {
	if url.Scheme!="http" && url.Scheme!="https" {
		panic("Protocol missing.")
	}
	if url.Host == "" {
		panic("Host missing.")
	}
	if url.Port() == "80" || url.Port() == "443" {
		log.Println("DO NOT add http's default port, it may affect the accuracy!")
	}
}