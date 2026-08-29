//go:build ignore

package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	serverCertPEM, err := os.ReadFile("server-cert.crt")
	if err != nil {
		log.Fatal(err)
	}

	clientCert, err := tls.LoadX509KeyPair("cert.crt", "client.key")
	if err != nil {
		log.Fatal(err)
	}

	rootCA := x509.NewCertPool()
	rootCA.AppendCertsFromPEM(serverCertPEM)

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      rootCA,
				Certificates: []tls.Certificate{clientCert},
			},
		},
	}

	resp, err := client.Get("https://localhost:8443")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response status:", resp.Status)
	log.Println("Response body:", string(body))
}
