# LevelGO

### Go Leveldb的RPC封装

GoLevelDB仓库[https://github.com/syndtr/goleveldb](https://github.com/syndtr/goleveldb)

使用方法:

```
go run cmd.go

go build
```
即可开启Leveldb rpc

![截图](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20201231214623105.png)

读性能

![读性能](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20201231214714901.png)

写性能

![写性能](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20201231214750151.png)

客户端库在external\client,有Get,Set,Has,Del,请参见对应的test

