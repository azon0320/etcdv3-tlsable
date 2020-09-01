package etcdv3tls

import (
	"os"
	"testing"
)

func TestEtcdSecure(t *testing.T) {
	os.Setenv(ENV_USETLS, "true")
	reg := NewRegistry()
	if !reg.Options().Secure {
		t.Fatalf("expect: %v, got: %v", true, reg.Options().Secure)
	}
}
