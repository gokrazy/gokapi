package gokapi

import (
	"context"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/gokrazy/gokapi/ondeviceapi"
	"github.com/gokrazy/gokrazy"
)

// ConnectOnDevice is meant to be run from a program running on a gokrazy
// instance and establishes a connection to the gokrazy on-device API.
func ConnectOnDevice() (*ondeviceapi.Configuration, error) {
	pw, err := os.ReadFile("/etc/gokr-pw.txt")
	if err != nil {
		return nil, err
	}
	cfg := ondeviceapi.NewConfiguration()
	auth := "gokrazy:" + strings.TrimSpace(string(pw))
	cfg.BasePath = "http://" + auth + "@unix/"

	// Use the Unix domain socket if available.
	if conn, err := net.Dial("unix", gokrazy.HTTPUnixSocket); err == nil {
		conn.Close()
		cfg.HTTPClient = &http.Client{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
					dialer := net.Dialer{}
					return dialer.DialContext(ctx, "unix", gokrazy.HTTPUnixSocket)
				},
			},
		}
		return cfg, nil
	}

	// Fallback to TCP.
	port, err := os.ReadFile("/etc/http-port.txt")
	if err != nil {
		return nil, err
	}
	cfg.BasePath = "http://" + auth + "@localhost:" + strings.TrimSpace(string(port)) + "/"
	return cfg, nil
}
