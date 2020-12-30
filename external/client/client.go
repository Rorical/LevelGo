package client

import (
	"LevelGo/external/utils"
	"LevelGo/internal/levelrpc"
	"context"
	"log"
	"strconv"

	grpc "google.golang.org/grpc"
)

type LevelRpcClient struct {
	Host   string
	Port   uint
	conn   *grpc.ClientConn
	Client *levelrpc.LevelRpcServiceClient
}

func RpcClient(host string, port uint) *LevelRpcClient {
	return &LevelRpcClient{
		Host: host,
		Port: port,
	}
}

func (self *LevelRpcClient) Connect() {
	port := strconv.FormatUint(uint64(self.Port), 10)
	conn, err := grpc.Dial(self.Host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connect Fail: %v", err)
	}
	self.conn = conn
	client := levelrpc.NewLevelRpcServiceClient(conn)
	self.Client = &client
}

func (self *LevelRpcClient) Get(key []byte) ([]byte, error) {
	reply, err := (*self.Client).Get(context.Background(), &levelrpc.GetRequest{Key: key})
	if err != nil {
		log.Fatalf("%v", err)
	}
	return reply.Value, utils.ErrCodeToError(reply.Error)

}

func (self *LevelRpcClient) Set(key []byte, value []byte) error {
	reply, err := (*self.Client).Set(context.Background(), &levelrpc.SetRequest{Key: key, Value: value})
	if err != nil {
		log.Fatalf("%v", err)
	}
	return utils.ErrCodeToError(reply.Error)
}

func (self *LevelRpcClient) Has(key []byte) (bool, error) {
	reply, err := (*self.Client).Has(context.Background(), &levelrpc.GetRequest{Key: key})
	if err != nil {
		log.Fatalf("%v", err)
	}
	return reply.Value, utils.ErrCodeToError(reply.Error)
}

func (self *LevelRpcClient) Del(key []byte) error {
	reply, err := (*self.Client).Del(context.Background(), &levelrpc.GetRequest{Key: key})
	if err != nil {
		log.Fatalf("%v", err)
	}
	return utils.ErrCodeToError(reply.Error)
}

func (self *LevelRpcClient) Close() {
	self.conn.Close()
}
