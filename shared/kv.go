package shared

import (
	"context"
	"github.com/hashicorp/go-plugin"
	"github.com/wulie/go-plugin-grpc/proto"
	"google.golang.org/grpc"
)

type KV interface {
	Put(key string, value []byte) error
	Get(key string) ([]byte, error)
}

type KVGRPCPlugin struct {
	plugin.Plugin
	Impl KV
}

func (k *KVGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterKVServer(s, &KVGRPCServer{Impl: k.Impl})
	return nil
}

func (k *KVGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, client *grpc.ClientConn) (interface{}, error) {
	return &KVGRPCClient{client: proto.NewKVClient(client)}, nil
}

var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

var PluginMap = map[string]plugin.Plugin{
	"kv_grpc": &KVGRPCPlugin{},
}
