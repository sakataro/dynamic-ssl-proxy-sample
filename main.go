package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	originURL := os.Getenv("ORIGIN_URL")
	origin, err := url.Parse(originURL)
	if err != nil {
		log.Fatal("invalid origin")
	}

	proxy := httputil.NewSingleHostReverseProxy(origin)

	config := &tls.Config{
		GetCertificate: getCert,
	}
	listner, err := tls.Listen("tcp", ":443", config)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.Serve(listner, proxy))
}

func getCert(info *tls.ClientHelloInfo) (*tls.Certificate, error) {

	// load key pair from file
	cer, err := tls.LoadX509KeyPair(
		"/shared/ssl/"+info.ServerName+".pem",
		"/shared/ssl/"+info.ServerName+"-key.pem",
	)

	// load key pair from redis

	if err != nil {
		return nil, err
	}
	return &cer, nil
}
