package main

import (
	"fmt"
	"net/http"

	"ajcode404.github.io/m/conn"
)

type CNamePool struct {
	pool *conn.CPool
	name string
}

type ShardPool struct {
	pool []*CNamePool
}

func getShardName(id int) string {
	return fmt.Sprintf("shard-%d", i)
}

func initShardConn() *ShardPool {
	shardPool := &ShardPool{
		pool: make([]*CNamePool, 0, 10),
	}
	for i := 0; i < 10; i++ {
		shardName := getShardName(i)
		cpool, err := conn.NewCPool(5, shardName)
		if err != nil {
			panic(err)
		}
		cNamePool := &CNamePool{cpool, shardName}
		shardPool.pool = append(shardPool.pool, cNamePool)
	}
	return shardPool
}

func (s *ShardPool) GetShard(id int) *CNamePool {
	shardId := id % 10
	shardName := getShardName(shardId)
	for i := range s.pool {
		if s.pool[i].name == shardName {
			return s.pool[i]
		}
	}
	return nil
}

// Implement Database Sharding and Routing (from API server)
func shard(w http.ResponseWriter, r *http.Request) {

	// get shards for depending on the ID's

}

func main() {
	// creating of shard brain connections

	http.HandleFunc("/shard", shard)
}
