package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"ajcode404.github.io/m/conn"
)

var (
	shardPoolInstance *ShardPool
	once              sync.Once
)

type CNamePool struct {
	pool *conn.CPool
	name string
}

type ShardPool struct {
	pool []*CNamePool
}

func getShardName(id int) string {
	return fmt.Sprintf("shard_%d", id)
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

func (s *ShardPool) GetShard(userId int) *CNamePool {
	shardId := userId % 10
	shardName := getShardName(shardId)
	for i := range s.pool {
		if s.pool[i].name == shardName {
			return s.pool[i]
		}
	}
	return nil
}

func GetShardPoolInstance() *ShardPool {
	once.Do(func() {
		shardPoolInstance = initShardConn()
	})
	return shardPoolInstance
}

// Implement Database Sharding and Routing (from API server)
func shard(w http.ResponseWriter, r *http.Request) {
	shardPool := GetShardPoolInstance()
	userId, err := strconv.Atoi(r.URL.Query().Get("userId"))
	if err != nil {
		fmt.Fprintln(w, "Provide the userID")
		return
	}
	shard := shardPool.GetShard(userId)
	conn, err := shard.pool.Get()
	if err != nil {
		panic(err)
	}
	conn.DB.Query("SELECT SLEEP(5.0)")
	if err != nil {
		panic(err)
	}
	shard.pool.Put(conn)
	fmt.Fprintf(w, "Write successful on %s\n", conn.DBName)
}

func main() {
	// creating of shard brain connections
	http.HandleFunc("/shard", shard)
	http.ListenAndServe(":8080", nil)
}
