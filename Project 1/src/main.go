package main

import (
  "fmt"
  "math/rand"
  "time"
	"container/list"
)

var ticks = 1000
var tick_duration = 1000000000
var C = 1
var L = 1000

func main() {

// SIMULATION 1
	queue := Queue{list.New()}
  simulator := Simulator{
		lambda:100,
		transmissionRate:C,
		packetSize:L,
		runTime:tick_duration,
		bufferSize:-1,
		queue:queue}

	simulator.initializeSimulator()

  for i := 0; i < 5; i++ {
    simulator.startSimulation()
		simulator.computeResults()
  }
	fmt.Printf("Avg packets in queue: %v\n", simulator.results.avgQueueSize)
}

// Random number generator
func randGenerator() (float64) {
  rand.Seed(time.Now().UTC().UnixNano())
	return rand.Float64()
}
