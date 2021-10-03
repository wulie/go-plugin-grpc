package main

import (
	"fmt"
	"github.com/hashicorp/go-plugin"
	"github.com/wulie/go-plugin-grpc/shared"
	"io/ioutil"
)

type KVImpl struct {
}

func (k *KVImpl) Put(key string, value []byte) error {
	value = []byte(fmt.Sprintf("%s\n\nWritten from plugin-go-grpc", string(value)))
	return ioutil.WriteFile("kv_"+key, value, 0644)
}

func (k *KVImpl) Get(key string) ([]byte, error) {
	return ioutil.ReadFile("kv_" + key)
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"kv": &shared.KVGRPCPlugin{Impl: &KVImpl{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
