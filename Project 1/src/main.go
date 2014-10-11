/*
	ECE 358 Project 1
	Christopher Poenaru		|		cpoenaru		|		20409287
	Paul Trautrim					|		pctrautr		|		--------

	This code is original and is the work of us as partners.
 */

package main

import (
  "fmt"
)

func main() {

// SIMULATION 1
	// Inputs
	ticks := 100000000
	tickDuration := 1000000
	C := 1
	L := 2000
	M := 5
	bufferSize := 5

	// Server instantiation
  simulator := Simulator{lambda:100, transmissionRate:C, packetSize:L,
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
		fmt.Printf(" ρ (Utilization of queue): \t%v\t|", results.finalQueueUtilization)
		fmt.Printf(" E[T] (Avg sojourn time): \t%v\n", results.finalAvgSojournTime)
	fmt.Printf("| Tick Duration:         \t%v\t|", simulator.tickDuration)
		fmt.Printf(" E[N] (Avg size of queue): \t%v\t|", results.finalAvgQueueSize)
		fmt.Printf(" Pidle (Time server is idle):\t%v\n", results.finalIdleTicks)
	fmt.Printf("| λ (Avg number of packets):\t%v\t|", simulator.transmissionRate)
		fmt.Printf(" Packets Received: \t%v\t\t|", results.finalReceivedPackets)
		fmt.Printf(" Ploss (Packet Loss Prob.):\t%v\n", results.finalLossProbability)
	fmt.Printf("| L (Length of packet [bits]):\t%v\t|", results.finalQueueUtilization)
		fmt.Printf(" Packets Dropped: \t%v\t\t|", results.finalDroppedPackets)
		fmt.Printf(" M (Simulation Iterations): \t%v\n", M)
	fmt.Printf("| C (Service time [ticks]):\t%v\t|", results.finalAvgQueueSize)
		fmt.Printf(" Packets Sent:     \t%v\t\t|", results.finalSentPackets)
		fmt.Printf("\t\t\t\t\t\t\n")
	fmt.Print("|-------------------------------------------------------------------------------------------------------------------------------|\n")
}
