package ssid

import (
	"fmt"
	"os/exec"
)

// Provider is an OS agnostic interface to fetch SSID information.
type Provider interface {
	// GetSSID retrieves the SSID for the specified device.
	GetSSID(device string) (string, error)
}

// execute an arbitrary command and return the output as string
func execute(cmd *exec.Cmd) (string, error) {
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("command execution error: %s", err.Error())
	}
	return string(out), nil
}
