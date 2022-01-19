package redis

import (
	goredis "github.com/go-redis/redis/v8"
)

type Settings struct {
	Default string   `yaml:"default"`
	Cluster []string `yaml:"cluster"`
}

type RdsClient struct {
	*goredis.Client
}

type RdsCluster struct {
	*goredis.ClusterClient
}

var (
	DB      *RdsClient
	Cluster *RdsCluster
)

func initialize(redisSetting *Settings) {
	if redisSetting.Default != "" {
		Cluster = nil
		opt, err := goredis.ParseURL(redisSetting.Default)
		if err != nil {
			panic(err)
		}

		redisClient := goredis.NewClient(opt)
		DB = &RdsClient{redisClient}
	} else {
		DB = nil
		clusterClient := goredis.NewClusterClient(&goredis.ClusterOptions{
			Addrs: redisSetting.Cluster,
		})
		Cluster = &RdsCluster{clusterClient}
	}
}
