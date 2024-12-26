package conn

import (
	"database/sql"
	"os"
	"sync"

	"github.com/go-sql-driver/mysql"
)

type Conn struct {
	DB *sql.DB
}

type CPool struct {
	mu      *sync.Mutex
	channel chan interface{}
	conns   []*Conn
	maxConn int
}

func NewCPool(maxConn int, dbName string) (*CPool, error) {
	var mu = sync.Mutex{}
	pool := &CPool{
		mu:      &mu,
		conns:   make([]*Conn, 0, maxConn),
		maxConn: maxConn,
		channel: make(chan interface{}, maxConn),
	}
	for i := 0; i < maxConn; i++ {
		pool.conns = append(pool.conns, &Conn{NewCon()})
		pool.channel <- struct{}{}
	}
	return pool, nil
}

func (pool *CPool) Close() {
	close(pool.channel)
	for i := range pool.conns {
		pool.conns[i].DB.Close()
	}
}

func (pool *CPool) Get() (*Conn, error) {
	<-pool.channel

	pool.mu.Lock()
	c := pool.conns[0]
	pool.conns = pool.conns[1:]
	pool.mu.Unlock()

	return c, nil

}

func (pool *CPool) Put(c *Conn) {
	pool.mu.Lock()
	pool.conns = append(pool.conns, c)
	pool.channel <- struct{}{}
	pool.mu.Unlock()
}

func NewCon() *sql.DB {

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	return db
}
