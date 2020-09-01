# etcdv3_tlsable
secure configurable etcdv3 go-micro plugin

# Compiling to micro
### with github.com/micro/micro/v2
~~~
// directory-structure
// - main.go
// - plugin.go

// file: plugin.go

package main

import _ "github.com/azon0320/etcdv3-tlsable"
~~~
### Then compile them
~~~
go build -o micro.executable main.go plugin.go
~~~

# Use registry in environment
~~~
MICRO_REGISTRY=etcdv3_tlsable \
MICRO_REGISTRY_ADDRESS=localhost:2379 \
MICRO_ETCDV3_USERNAME=username \
MICRO_ETCDV3_PASSWORD=password \
MICRO_ETCDV3_SECURE=true \
micro web
~~~

# etcdv3_tlsable Environments
## MICRO_ETCDV3_USERNAME
etcd user name, plugin will not create credential(username,password) if it empty.
### Example
~~~
MICRO_ETCDV3_USERNAME=thename micro web
~~~

## MICRO_ETCDV3_PASSWORD
etcd password, plugin will ignore if `MICRO_ETCDV3_USERNAME` empty.
### Example
~~~
MICRO_ETCDV3_USERNAME=thename MICRO_ETCDV3_PASSWORD=thepasswd micro web
~~~

## MICRO_ETCDV3_SECURE
Specify if enable secure connection
### Example
~~~
MICRO_ETCDV3_SECURE=true micro web
~~~
