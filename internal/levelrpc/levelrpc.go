package levelrpc

import (
	context "context"
	"log"
	"net"
	"strconv"

	"LevelGo/external/utils"
	"LevelGo/internal/config"
	"LevelGo/internal/leveldb"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type LevelServer struct {
	DB *leveldb.LevelDB
}

func (s *LevelServer) Get(ctx context.Context, in *GetRequest) (*GetReply, error) {
	log.Printf("Get %s", in.Key)
	value, err := s.DB.Get(in.Key)
	return &GetReply{Value: value, Error: utils.ErrorToErrCode(err)}, nil
}

func (s *LevelServer) Set(ctx context.Context, in *SetRequest) (*ErrorReply, error) {
	log.Printf("Set %s", in.Key)
	err := s.DB.Set(in.Key, in.Value)
	return &ErrorReply{Error: utils.ErrorToErrCode(err)}, nil
}

func (s *LevelServer) Has(ctx context.Context, in *GetRequest) (*HasReply, error) {
	log.Printf("Has %s", in.Key)
	value, err := s.DB.Has(in.Key)
	return &HasReply{Value: value, Error: utils.ErrorToErrCode(err)}, nil
}

func (s *LevelServer) Del(ctx context.Context, in *GetRequest) (*ErrorReply, error) {
	log.Printf("Del %s", in.Key)
	err := s.DB.Del(in.Key)
	return &ErrorReply{Error: utils.ErrorToErrCode(err)}, nil
}

type LevelRpcServer struct {
	Port uint
	DB   *leveldb.LevelDB
}

func RpcServer(leveldbConf *config.LevelDBSetting, db *leveldb.LevelDB) *LevelRpcServer {
	return &LevelRpcServer{
		Port: leveldbConf.RpcPort,
		DB:   db,
	}
}

func (self *LevelRpcServer) Listen() {
	var err error
	port := strconv.FormatUint(uint64(self.Port), 10)
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	RegisterLevelRpcServiceServer(server, &LevelServer{DB: self.DB})
	reflection.Register(server)
	log.Printf("RPC Listen at :%s", port)
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}

}
