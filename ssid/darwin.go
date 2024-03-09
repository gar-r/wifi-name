package ssid

import (
	"errors"
	"os/exec"
	"regexp"
)

var darwinRegex = regexp.MustCompile(`Current Wi-Fi Network: (.*)`)

type Darwin struct {
	CmdProvider
}

func (d *Darwin) GetSSID(device string) (string, error) {
	d.cmd = exec.Command("networksetup", "-getairportnetwork", device)
	out, err := d.CmdProvider.GetSSID(device)
	if err != nil {
		return "", err
	}
	matches := darwinRegex.FindStringSubmatch(out)
	if len(matches) > 1 {
		return matches[1], nil
	}
	return "", errors.New("ssid cannot be determined")
}
