package main

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestH2C(t *testing.T) {
	h2s := &http2.Server{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %v, http: %v", r.URL.Path, r.TLS == nil)
	})
	server := &http.Server{Addr: "0.0.0.0:10100", Handler: h2c.NewHandler(handler, h2s)}

	fmt.Printf("Listening [0.0.0.0:10100]...\n")
	if err := server.ListenAndServe(); err != nil {
		t.Fatal(err)
	}
}

func TestH2COnly2(t *testing.T) {
	server := http2.Server{}

	l, err := net.Listen("tcp", "0.0.0.0:1010")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Listening [0.0.0.0:1010]...\n")
	for {
		conn, err := l.Accept()
		if err != nil {
			t.Fatal(err)
		}

		server.ServeConn(conn, &http2.ServeConnOpts{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello, %v, http: %v", r.URL.Path, r.TLS == nil)
			}),
		})
	}
}

func TestH2CClient(t *testing.T) {
	clt := http.Client{
		Transport: &http2.Transport{
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
			AllowHTTP: true,
		},

		Timeout: time.Second * 10,
	}

	resp, err := clt.Get("https://tc:10100")
	if err != nil {
		t.Fatal(err)
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bs))
}

func TestH2WithTLS(t *testing.T) {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example.com"),
		Cache:      autocert.DirCache("certs"),
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	server := &http.Server{
		Addr: ":443",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	//http2.ConfigureServer(server, &http2.Server{})
	go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
	log.Fatal(server.ListenAndServeTLS("", ""))
}
