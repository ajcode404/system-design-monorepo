package sync

// count prime numbers
import (
	"fmt"
	"math"
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
	totalPrimeNo++
}

func Sync() {
	start := time.Now()
	for i := 3; i < MAX_INT; i++ {
		checkPrime(i)
	}
	fmt.Println("checking till", MAX_INT, "found", totalPrimeNo+1, "prime number took", time.Since(start))
}
