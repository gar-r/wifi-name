package ssid

import (
	"errors"
	"os/exec"
	"regexp"
)

var windowsRegex = regexp.MustCompile(`\s+SSID\s+:\s+(.*)`)

// Windows implements Provider on Windows.
type Windows struct{}

func (w *Windows) GetSSID(device string) (string, error) {
	cmd := exec.Command("cmd", "/c", "netsh", "wlan", "show", "interfaces")
	out, err := execute(cmd)
	if err != nil {
		return "", err
	}
	matches := windowsRegex.FindStringSubmatch(out)
	if len(matches) > 1 {
		return matches[1], nil
	}
	return "", errors.New("ssid cannot be determined")
}
