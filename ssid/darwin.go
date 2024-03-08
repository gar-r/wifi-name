package ssid

import (
	"os/exec"
	"strings"
)

const OutputPrefix = "Current Wi-Fi Network: "

type Darwin struct {
	CmdProvider
}

func (d *Darwin) GetSSID() (string, error) {
	d.cmd = exec.Command("networksetup", "-getairportnetwork", "en0")
	out, err := d.CmdProvider.GetSSID()
	if err != nil {
		return "", err
	}
	return strings.Replace(out, OutputPrefix, "", 1), nil
}
