package main

import (
	"math/rand"
	"time"
)

// Random number generator
func randGenerator() (float64) {
rand.Seed(time.Now().UTC().UnixNano())
return rand.Float64()
}
