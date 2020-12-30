package client

import (
	"LevelGo/external/utils"
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	client := RpcClient("localhost", 8192)
	client.Connect()
	h, err := client.Get(utils.StringIn("a"))
	fmt.Println(utils.StringOut(h))
	if err != nil {
		panic(err)
	}
	err = client.Set(utils.StringIn("a"), utils.StringIn("你好"))
	if err != nil {
		panic(err)
	}
	defer client.Close()
}
