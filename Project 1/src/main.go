package main

import (
//  "fmt"
  "math/rand"
  "time"
	"container/list"
	"fmt"
)

var ticks = 1000
var tick_duration = 100
var transmissionRate = 1

func main() {

  // Question 1
//  var lambda = 100
//  var C = 1
//  var L = 2000

	queue := Queue{list.New()}

  simulator := Simulator{queue:queue}
  simulator.startSimulation()
//	var M = 0

	queue.enqueue(Packet{remainingBits:9000})
	fmt.Println(queue.buffer.Len())
	fmt.Println(queue.dequeue().remainingBits)
	fmt.Println(queue.buffer.Len())



  for i := 0; i < ticks; i++ {
    generatePacket()
    servicePacket()
  }
}

// Random number generator
func randGenerator() (float64) {
  rand.Seed(time.Now().UTC().UnixNano())
	return rand.Float64()
}

// The Packet generator
func generatePacket() {
  //fmt.Println("Checking if packet needs to be generated")
}

// The server for the packets
func servicePacket() {
  //fmt.Println("Servicing packet")
}
