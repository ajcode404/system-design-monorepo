package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"ajcode404.github.io/m/conn"
)

func benchamrkPool() {
	startTime := time.Now()
	pool, err := conn.NewCPool(10)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			c, err := pool.Get()
			if err != nil {
				fmt.Printf("error is %s\n", err)
				panic(err)
			}
			_, dErr := c.DB.Exec("SELECT SLEEP(0.01)")
			if dErr != nil {
				panic(dErr)
			}
			pool.Put(c)
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

			db := conn.NewCon()

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
func main() {

	// benchamrkNoPool()
	benchamrkPool()
}
