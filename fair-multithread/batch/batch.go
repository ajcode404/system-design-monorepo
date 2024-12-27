package batch

// count prime numbers
import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var MAX_INT = 100_000_000
var TOTAL_THREADS = 10
var totalPrimeNo int32 = 0

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

func doBatch(threadName string, wg *sync.WaitGroup, nstart int, nend int) {
	defer wg.Done()
	start := time.Now()
	for i := nstart; i < nend; i++ {
		checkPrime(i)
	}
	fmt.Printf("thread %s [%d, %d] completed in %s\n", threadName, nstart, nend, time.Since(start))
}

func checkPrimeAsync() {
	var wg sync.WaitGroup
	nstart := 3
	batchSize := int(float64(MAX_INT) / float64(TOTAL_THREADS))
	start := time.Now()
	for i := 0; i < TOTAL_THREADS-1; i++ {
		wg.Add(1)
		go doBatch(strconv.Itoa(i), &wg, nstart, nstart+batchSize)
		nstart += batchSize
	}
	wg.Add(1)
	go doBatch(strconv.Itoa(TOTAL_THREADS-1), &wg, nstart, MAX_INT)
	wg.Wait()
	fmt.Printf("prime numbers %d in %d time\n", totalPrimeNo, time.Since(start))
}

func Batch() {
	checkPrimeAsync()
}
