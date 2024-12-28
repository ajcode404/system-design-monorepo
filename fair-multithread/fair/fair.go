package fair

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

var MAX_INT = 100_000_000
var TOTAL_THREADS = 10
var totalPrimeNo int32 = 0
var nextNumber int32 = 2

func checkPrime(x int) {
	// if number even skip
	if x&1 == 0 {
		return
	}
	for i := 3; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return
		}
	}
	atomic.AddInt32(&totalPrimeNo, 1)
}

func workerThreads(wg *sync.WaitGroup, threadNo int) {
	defer wg.Done()
	start := time.Now()
	no := atomic.AddInt32(&nextNumber, 1)
	for no < int32(MAX_INT) {
		checkPrime(int(no))
		no = atomic.AddInt32(&nextNumber, 1)
	}
	fmt.Printf("thread No %d took %s time\n", threadNo, time.Since(start))
}

func Fair() {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < TOTAL_THREADS; i++ {
		wg.Add(1)
		go workerThreads(&wg, i)
	}
	wg.Wait()
	fmt.Printf("prime no := %d time took := %s\n", totalPrimeNo, time.Since(start))
}
