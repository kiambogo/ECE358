package main

import (
  "fmt"
  "math/rand"
  "time"
	"container/list"
)

var ticks = 1000
var tick_duration = 100000
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
	fmt.Printf("# packets received: %v\n", simulator.results.packetsReceived)
	fmt.Printf("Summed sojurn time: %v\n", simulator.results.summedSojurnTime)

	fmt.Printf("Avg sojurn time: %v\n", simulator.results.avgSojurnTime)

	/*queue.enqueue(&Packet{L, 0})
	fmt.Printf("Remaining: %v\n", queue.peek().remainingBits)
	queue.peek().decrementRemainingBits(C)
	fmt.Printf("Remaining: %v\n", queue.peek().remainingBits)*/


}

// Random number generator
func randGenerator() (float64) {
  rand.Seed(time.Now().UTC().UnixNano())
	return rand.Float64()
}
