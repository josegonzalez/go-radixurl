package radixurl_test

import (
	"github.com/fzzy/radix/redis"
	"github.com/josegonzalez/go-radixurl"
	"testing"
)

func TestConnect(t *testing.T) {
	c, err := radixurl.Connect()

	if err != nil {
		t.Errorf("Error returned")
	}
}
