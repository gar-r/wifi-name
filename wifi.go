package wifiname

import (
	"fmt"
	"runtime"

	"git.okki.hu/garrichs/wifi-name/ssid"
)

var os = runtime.GOOS

func GetSSID(device string) (string, error) {
	p, err := getProvider()
	if err != nil {
		return "", err
	}
	return p.GetSSID(device)
}

func getProvider() (ssid.Provider, error) {
	switch os {
	case "darwin":
		return &ssid.Darwin{}, nil
	default:
		return nil, fmt.Errorf("unsupported os: '%s'", os)
	}
}
