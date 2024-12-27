package mod

// count prime numbers
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

func doBatch(modFactor int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	for i := 3; i < MAX_INT; i++ {
		if i%10 == modFactor {
			checkPrime(i)
		}
	}
	fmt.Printf("thread %d completed in %s\n", modFactor, time.Since(start))
}

func checkPrimeAsync() {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 1; i < TOTAL_THREADS; i++ {
		wg.Add(1)
		go doBatch(i, &wg)
	}
	wg.Wait()
	fmt.Printf("prime numbers %d in %s time\n", totalPrimeNo, time.Since(start))
}

func Mod() {
	checkPrimeAsync()
}
