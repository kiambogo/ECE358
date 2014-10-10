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
	tickDuration := 1000000
	C := 1
	L := 2000
	M := 5
	bufferSize := 10

	// Server instantiation
  simulator := Simulator{lambda:100, transmissionRate:C, packetSize:L, runTime:tickDuration, bufferSize:bufferSize}
	results := Results{}
  for i := 0; i < M; i++ {
		simulator.initializeSimulator()
    simulator.startSimulation()
		simulator.computeResults()
		results.updateResultSet(&simulator)
  }
	results.computeFinalResults(M)

	//Output + formatting
	fmt.Print("\n=================================================================================\n")
	fmt.Printf("|\t\t\t\tSIMULATION 1 RESULTS -- %v/D/%v \t\t\t|", M, 1)
	fmt.Print("\n|-------------------------------------------------------------------------------|\n")
	fmt.Printf("| λ (Avg number of packets):\t%v\t|", simulator.lambda)
		fmt.Printf(" E[T] (Avg sojourn time): \t%v\t|\n", results.finalAvgSojournTime)
	fmt.Printf("| L (Length of packet [bits]):\t%v\t|", simulator.packetSize)
		fmt.Printf(" Pidle (Time server is idle):\t%v\t|\n", results.finalIdleTicks)
	fmt.Printf("| C (Service time [ticks]): \t%v\t|", simulator.transmissionRate)
		fmt.Printf(" Ploss (Packet Loss Prob.):\t%v\t|\n", results.finalLossProbability)
	fmt.Printf("| ρ (Utilization of queue): \t%v\t|", results.finalQueueUtilization)
		fmt.Printf(" M (Simulation Iterations): \t%v\t|\n", M)
	fmt.Printf("| E[N] (Avg size of queue): \t%v\t|", results.finalAvgQueueSize)
		fmt.Printf("\t\t\t\t\t|\n")
	fmt.Print("|-------------------------------------------------------------------------------|\n")
}
