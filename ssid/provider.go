package ssid

import (
	"os/exec"
	"strings"
)

type Provider interface {
	GetSSID() (string, error)
}

type CmdProvider struct {
	cmd *exec.Cmd
}

func (c *CmdProvider) GetSSID() (string, error) {
	out, err := c.cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
