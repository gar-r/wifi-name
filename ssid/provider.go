package ssid

import (
	"os/exec"
	"strings"
)

type Provider interface {
	GetSSID(device string) (string, error)
}

type CmdProvider struct {
	cmd *exec.Cmd
}

func (c *CmdProvider) GetSSID(_ string) (string, error) {
	out, err := c.cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
