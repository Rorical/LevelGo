package client

import (
	"LevelGo/external/utils"
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	client := RpcClient("localhost", 8192)
	client.Connect()
	h, err := client.Get(utils.StringIn("15893"))
	t.Log(utils.StringOut(h))
	if err != nil {
		panic(err)
	}
	defer client.Close()
}

func BenchmarkClient(b *testing.B) {
	client := RpcClient("localhost", 8192)
	client.Connect()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.Get(utils.StringIn(fmt.Sprint(i)))
	}
	defer client.Close()
}
