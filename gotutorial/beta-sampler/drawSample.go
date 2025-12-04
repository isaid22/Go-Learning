// main.go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	// Seed a RNG (use your own seed if you want reproducible results)
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	// Beta(α=2, β=5)
	b := distuv.Beta{
		Alpha: 2,
		Beta:  5,
		Src:   rng, // optional; if nil, uses global math/rand
	}

	// Draw a few samples
	for i := 0; i < 5; i++ {
		fmt.Printf("%.6f\n", b.Rand())
	}

	// You can also access PDF/Mean/Var if needed
	fmt.Printf("Mean=%.4f Var=%.4f PDF(0.3)=%.4f\n", b.Mean(), b.Variance(), b.Prob(0.3))
}
