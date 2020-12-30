package levelgo

import (
	"LevelGo/internal/config"
	"LevelGo/internal/leveldb"
	"LevelGo/internal/levelrpc"
)

func Run() {
	conf := config.Read()
	level := leveldb.GetLevelDB(conf)
	rpcSer := levelrpc.RpcServer(conf, level)
	rpcSer.Listen()
}
