package etcdv3tls

import (
	"fmt"
	"os"

	"github.com/micro/go-micro/v2/config/cmd"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

const (
	ENV_DEBUG    = "MICRO_ETCDV3_DEBUG"
	ENV_USERNAME = "MICRO_ETCDV3_USERNAME"
	ENV_PASSWORD = "MICRO_ETCDV3_PASSWORD"
	ENV_SECURE   = "MICRO_ETCDV3_SECURE"
)

func init() {
	cmd.DefaultRegistries["etcdv3_tlsable"] = NewRegistry
}

func isDebug() bool {
	return os.Getenv(ENV_DEBUG) == "true"
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	presetOpts := make([]registry.Option, 0)
	if os.Getenv(ENV_USERNAME) != "" {
		if isDebug() {
			fmt.Println("[ETCDV3TLS] inject credentials")
		}
		presetOpts = append(
			presetOpts,
			etcd.Auth(os.Getenv(ENV_USERNAME), os.Getenv(ENV_PASSWORD)),
		)
	}
	if os.Getenv(ENV_SECURE) == "true" {
		if isDebug() {
			fmt.Println("[ETCDV3TLS] enable secure")
		}
		presetOpts = append(
			presetOpts,
			registry.Secure(true),
		)
	}
	presetOpts = append(presetOpts, opts...)
	return etcd.NewRegistry(presetOpts...)
}
