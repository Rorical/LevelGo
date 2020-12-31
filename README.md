# LevelGO

### Go Leveldb的RPC封装

GoLevelDB仓库[https://github.com/syndtr/goleveldb](https://github.com/syndtr/goleveldb)

使用方法:

```
go run cmd.go

go build
```
即可开启Leveldb rpc

![截图](https://doc.kmf.com/ke-feedback/2020/12/31/21/46/18/20201231214617.png)

读性能

![读性能](https://doc.kmf.com/ke-feedback/2020/12/31/21/47/11/20201231214711.png)

写性能

![写性能](https://doc.kmf.com/ke-feedback/2020/12/31/21/47/45/M59]S5{G6~BVSS2NPV{Z]YI.png)

客户端库在external\client,有Get,Set,Has,Del,请参见对应的test

