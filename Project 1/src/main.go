/*
	ECE 358 Project 1
	Christopher Poenaru  |  cpoenaru  |  20409287
	Paul Trautrim        |  pctrautr  |  20348861

	This code is original and is the work of us as partners.
 */

package main

import (
  "fmt"
	"time"
)

func main() {

	var ticks, tickDuration uint64
	var lambda float32
	var C, L uint32
	var M, bufferSize int

// Question 2
	// Inputs
	ticks = uint64(60000000000)
	tickDuration = uint64(1000000000)
	lambda = 100
	C = 1
	L = 2000000
	M = 5
	bufferSize = -1

	// Server instantiation
	startTime := time.Now()
  simulator := Simulator{lambda:lambda, transmissionRate:C, packetSize:L,
		bufferSize:bufferSize, tickDuration:tickDuration, ticks:ticks}
	results := Results{}
  for i := 0; i < M; i++ {
    simulator.startSimulation()
		simulator.computeResults()
		results.updateResultSet(&simulator)
  }
	results.computeFinalResults(M)

	//Output + formatting
	fmt.Print("\n=================================================================================================================================\n")
	fmt.Printf("|\t\t\t\t\t\tSIMULATION 1 RESULTS -- %v/D/%v \t\t\t\t\t\t\t|", M, 1)
	fmt.Print("\n|-------------------------------------------------------------------------------------------------------------------------------|\n")
	fmt.Printf("| Number of Ticks:    \t%v\t|", simulator.ticks)
		fmt.Printf(" ρ (Queue utilization): \t%v\t|", float32(float32(L) * (lambda/float32(C))))
		fmt.Printf(" E[T] (Avg sojourn time): \t%v\n", float32(results.finalAvgSojournTime))
	fmt.Printf("| Tick Duration:      \t%v\t|", simulator.tickDuration)
		fmt.Printf(" E[N] (Avg size of queue): \t%v\t|", float32(results.finalAvgQueueSize))
		fmt.Printf(" Pidle (Time server is idle):\t%v\n", float32(results.finalIdleServerProp))
	fmt.Printf("| λ (Avg number of packets):\t%v\t|", simulator.lambda)
		fmt.Printf(" Packets Received: \t%v\t\t|", results.finalReceivedPackets)
		fmt.Printf(" Ploss (Packet Loss Prob.):\t%v\n", results.finalLossProbability)
	fmt.Printf("| L (Length of packet [bits]):\t%v\t|", simulator.packetSize)
		fmt.Printf(" Packets Dropped: \t%v\t\t|", results.finalDroppedPackets)
		fmt.Printf(" M (Simulation Iterations): \t%v\n", M)
	fmt.Printf("| C (Service time [ticks]):\t%v\t|", simulator.transmissionRate)
		fmt.Printf(" Packets Sent:     \t%v\t\t|", results.finalSentPackets)
		fmt.Printf(" Runtime Duration: \t%v\n", (time.Now().Sub(startTime)))
	fmt.Print("|-------------------------------------------------------------------------------------------------------------------------------|\n")

	// Question 3
	// Inputs
	ticks = uint64(60000000000)
	tickDuration = uint64(1000000000)
	C = 1
	L = 2000000
	M = 5
	bufferSize = -1

	// Server instantiation

	for i := 0.2; i < 0.9; i+=0.1 {
		startTime := time.Now()
		lambda = float32((i * float64(C) * float64(tickDuration)) / float64(L))
		simulator := Simulator{lambda:lambda, transmissionRate:C, packetSize:L,
			bufferSize:bufferSize, tickDuration:tickDuration, ticks:ticks}
		results := Results{}

		simulator.startSimulation()
		simulator.computeResults()
		results.updateResultSet(&simulator)
		results.computeFinalResults(M)

		//Output + formatting
		fmt.Print("\n=================================================================================================================================\n")
		fmt.Printf("|\t\t\t\t\t\tSIMULATION 2 (ρ=%v) RESULTS -- %v/D/%v \t\t\t\t\t\t\t|", i, M, 1)
		fmt.Print("\n|-------------------------------------------------------------------------------------------------------------------------------|\n")
		fmt.Printf("| Number of Ticks:    \t%v\t|", simulator.ticks)
		fmt.Printf(" ρ (Queue utilization): \t%v\t|", i)
		fmt.Printf(" E[T] (Avg sojourn time): \t%v\n", float32(results.finalAvgSojournTime))
		fmt.Printf("| Tick Duration:      \t%v\t|", simulator.tickDuration)
		fmt.Printf(" E[N] (Avg size of queue): \t%v\t|", float32(results.finalAvgQueueSize))
		fmt.Printf(" Pidle (Time server is idle):\t%v\n", float32(results.finalIdleServerProp))
		fmt.Printf("| λ (Avg number of packets):\t%v\t|", simulator.lambda)
		fmt.Printf(" Packets Received: \t%v\t\t|", results.finalReceivedPackets)
		fmt.Printf(" Ploss (Packet Loss Prob.):\t%v\n", results.finalLossProbability)
		fmt.Printf("| L (Length of packet [bits]):\t%v\t|", simulator.packetSize)
		fmt.Printf(" Packets Dropped: \t%v\t\t|", results.finalDroppedPackets)
		fmt.Printf(" M (Simulation Iterations): \t%v\n", M)
		fmt.Printf("| C (Service time [ticks]):\t%v\t|", simulator.transmissionRate)
		fmt.Printf(" Packets Sent:     \t%v\t\t|", results.finalSentPackets)
		fmt.Printf(" Runtime Duration: \t%v\n", (time.Now().Sub(startTime)))
		fmt.Print("|-------------------------------------------------------------------------------------------------------------------------------|\n")
	}
}
