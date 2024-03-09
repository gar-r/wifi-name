package ssid

import (
	"errors"
	"os/exec"
)

type Linux struct{}

func (l *Linux) GetSSID(device string) (string, error) {
	cmd := exec.Command("iwgetid", "-r", device)
	out, err := execute(cmd)
	if err != nil {
		return "", err
	}
	if out == "" {
		return "", errors.New("ssid cannot be determined")
	}
	return out, nil
}
