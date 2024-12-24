package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
)

type conn struct {
	db *sql.DB
}

type cpool struct {
	mu      *sync.Mutex
	channel chan interface{}
	conns   []*conn
	maxConn int
}

func NewCPool(maxConn int) (*cpool, error) {
	var mu = sync.Mutex{}
	pool := &cpool{
		mu:      &mu,
		conns:   make([]*conn, 0, maxConn),
		maxConn: maxConn,
		channel: make(chan interface{}, maxConn),
	}
	for i := 0; i < maxConn; i++ {
		pool.conns = append(pool.conns, &conn{newCon()})
		pool.channel <- struct{}{}
	}
	return pool, nil
}

func (pool *cpool) Close() {
	close(pool.channel)
	for i := range pool.conns {
		pool.conns[i].db.Close()
	}
}

func (pool *cpool) Get() (*conn, error) {
	<-pool.channel

	pool.mu.Lock()
	c := pool.conns[0]
	pool.conns = pool.conns[1:]
	pool.mu.Unlock()

	return c, nil

}

func (pool *cpool) Put(c *conn) {
	pool.mu.Lock()
	pool.conns = append(pool.conns, c)
	pool.channel <- struct{}{}
	pool.mu.Unlock()
}

func benchamrkPool() {
	startTime := time.Now()
	pool, err := NewCPool(10)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			conn, err := pool.Get()
			if err != nil {
				fmt.Printf("error is %s\n", err)
				panic(err)
			}
			_, dErr := conn.db.Exec("SELECT SLEEP(0.01)")
			if dErr != nil {
				panic(dErr)
			}
			pool.Put(conn)
		}()
	}

	wg.Wait()
	pool.Close()
	log.Printf("Took %s time\n", time.Since(startTime))
}

func benchamrkNoPool() {
	startTime := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			db := newCon()

			_, err := db.Exec("SELECT SLEEP(0.01)")
			if err != nil {
				panic(err)
			}
			db.Close()
		}()
	}
	wg.Wait()
	log.Printf("Took %s time\n", time.Since(startTime))
}

func newCon() *sql.DB {

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

func main() {

	// benchamrkNoPool()
	benchamrkPool()
}
