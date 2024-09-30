package gokapi

import (
	"os"

	"github.com/gokrazy/internal/config"
	"github.com/gokrazy/internal/instanceflag"
	"github.com/gokrazy/internal/updateflag"
)

func BasePathForInstanceFlag() (string, error) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		if os.IsNotExist(err) {
			// best-effort compatibility for old setups
			cfg = config.NewStruct(instanceflag.Instance())
		} else {
			return "", err
		}
	}

	updateflag.SetUpdate("yes")

	update, err := cfg.Update.WithFallbackToHostSpecific(cfg.Hostname)
	if err != nil {
		return "", err
	}

	if update.HTTPPort == "" {
		update.HTTPPort = "80"
	}

	if update.HTTPSPort == "" {
		update.HTTPSPort = "443"
	}

	// TODO: support for TLS certificates
	const schema = "http"
	u, err := updateflag.BaseURL(update.HTTPPort, update.HTTPSPort, schema, update.Hostname, update.HTTPPassword)
	if err != nil {
		return "", err
	}

	return u.String(), nil
}
