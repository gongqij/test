package redis

import (
	"context"
	"testing"
	"time"
)

var (
	MASTER_NAME       = "mymaster"
	SENTINEL_ADDRS    = []string{"192.168.10.45:36378", "192.168.10.46:36378", "192.168.10.47:36378"}
	SENTINEL_PASSWORD = ""
	SENTINEL_DB       = 1
)

func TestRedis(t *testing.T) {
	initialize(&Settings{
		Default: "redis://127.0.0.1:6379",
		Cluster: nil,
	})
	for {
		time.Sleep(time.Second)
		DB.Set(context.Background(), "gq", "hahahahah", time.Hour)
	}
}

func TestNewRedisSentinel(t *testing.T) {
	NewRedisSentinel(MASTER_NAME, SENTINEL_PASSWORD, SENTINEL_ADDRS)
}
