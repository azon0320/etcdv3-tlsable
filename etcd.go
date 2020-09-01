package etcdv3tls

import (
	"fmt"
	"os"
	"strings"

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
	/*
		go func() {
			debugR := mux.NewRouter()
			debugR.HandleFunc("/debug/pprof/", pprof.Index)
			debugR.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
			debugR.HandleFunc("/debug/pprof/profile", pprof.Profile)
			debugR.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
			debugR.Handle("/debug/pprof/heap", pprof.Handler("heap"))
			debugR.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
			debugR.Handle("/debug/pprof/block", pprof.Handler("block"))
			debugR.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
			err := http.ListenAndServe(":9999", debugR)
			if err != nil {
				fmt.Print(err)
			}
		}()*/
}

func isDebug() bool {
	return os.Getenv(ENV_DEBUG) == "true"
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	if isDebug() {
		mockOpts := registry.Options{}
		for _, opt := range opts {
			opt(&mockOpts)
		}
		fmt.Println("[ETCDV3TLS] preset addrs ", strings.Join(mockOpts.Addrs, ","))
		fmt.Println("[ETCDV3TLS] preset secure ", mockOpts.Secure)
		fmt.Println("[ETCDV3TLS] preset timeout ", mockOpts.Timeout)
	}
	presetOpts := make([]registry.Option, 0)
	if os.Getenv(ENV_USERNAME) != "" {
		if isDebug() {
			fmt.Println("[ETCDV3TLS] inject auth")
		}
		username := os.Getenv(ENV_USERNAME)
		if isDebug() {
			fmt.Println("[ETCDV3TLS] attempt got username from env")
		}
		password := os.Getenv(ENV_PASSWORD)
		if isDebug() {
			fmt.Println("[ETCDV3TLS] attempt got password from env")
		}
		auth := etcd.Auth(username, password)
		if isDebug() {
			fmt.Println("[ETCDV3TLS] auth option generated")
		}
		presetOpts = append(
			presetOpts,
			auth,
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
	presetOpts = append(opts, presetOpts...)
	return etcd.NewRegistry(presetOpts...)
}
