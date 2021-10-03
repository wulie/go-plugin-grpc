package shared

import (
	"context"
	"github.com/wulie/go-plugin-grpc/proto"
)

type KVGRPCClient struct {
	client proto.KVClient
}

func (k *KVGRPCClient) Put(key string, value []byte) error {
	_, err := k.client.Put(context.Background(), &proto.PutRequest{
		Key:   key,
		Value: value,
	})
	return err
}

func (k *KVGRPCClient) Get(key string) ([]byte, error) {
	rep, err := k.client.Get(context.Background(), &proto.GetRequest{
		Key: key,
	})
	if err != nil {
		return nil, err
	}
	return rep.Value, nil
}

type KVGRPCServer struct {
	Impl KV
}

func (k *KVGRPCServer) Get(ctx context.Context, request *proto.GetRequest) (*proto.GetResponse, error) {
	bytes, err := k.Impl.Get(request.Key)
	if err != nil {
		return &proto.GetResponse{}, err
	}
	return &proto.GetResponse{Value: bytes}, nil
}

func (k *KVGRPCServer) Put(ctx context.Context, request *proto.PutRequest) (*proto.Empty, error) {
	err := k.Impl.Put(request.Key, request.Value)
	return &proto.Empty{}, err
}
