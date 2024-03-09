package ssid

import (
	"errors"
	"os/exec"
	"regexp"
)

var darwinRegex = regexp.MustCompile(`Current Wi-Fi Network: (.*)`)

// Darwin implements Provider on Mac.
type Darwin struct{}

func (d *Darwin) GetSSID(device string) (string, error) {
	cmd := exec.Command("networksetup", "-getairportnetwork", device)
	out, err := execute(cmd)
	if err != nil {
		return "", err
	}
	matches := darwinRegex.FindStringSubmatch(string(out))
	if len(matches) > 1 {
		return matches[1], nil
	}
	return "", errors.New("ssid cannot be determined")
}
