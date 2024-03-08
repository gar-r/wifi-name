package wifiname

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSSID(t *testing.T) {

	t.Run("unsupported os", func(t *testing.T) {
		os = "foobar"
		defer func() {
			os = runtime.GOOS
		}()
		_, err := GetSSID()
		assert.Error(t, err)
	})

	t.Run("get ssid", func(t *testing.T) {
		ssid, err := GetSSID()
		assert.NoError(t, err)
		t.Logf("ssid: %s\n", ssid)
	})

}
