package main

import (
	"math/rand"
	"time"
)

// Random number generator
func randGenerator() (float64) {
//rand.Seed(time.Now().UTC().UnixNano())
//return rand.Float64()
	return rand.New(rand.NewSource(time.Now().UnixNano())).Float64()

}

