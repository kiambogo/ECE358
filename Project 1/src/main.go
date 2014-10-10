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
	tick_duration := 1000000
	C := 1
	L := 2000
	M := 5

	// Server instantiation
  simulator := Simulator{lambda:100, transmissionRate:C, packetSize:L, runTime:tick_duration, bufferSize:-1}
	simulator.initializeSimulator()

  for i := 0; i < M; i++ {
    simulator.startSimulation()
		simulator.computeResults()
  }

	//Output + formatting
	fmt.Print("\n=================================================================================\n")
	fmt.Printf("|\t\t\t\tSIMULATION 1 RESULTS -- %v/D/%v \t\t\t|", M, 1)
	fmt.Print("\n|-------------------------------------------------------------------------------|\n")
	fmt.Printf("| λ (Avg number of packets):\t%v\t|", simulator.lambda)
		fmt.Printf(" E[T] (Avg sojourn time): \t%v\t|\n", simulator.results.avgSojurnTime)
	fmt.Printf("| L (Length of packet [bits]):\t%v\t|", simulator.packetSize)
		fmt.Printf(" Pidle (Time server is idle):\t%v\t|\n", simulator.results.idleTicks)
	fmt.Printf("| C (Service time [ticks]): \t%v\t|", simulator.transmissionRate)
		fmt.Printf(" Ploss (Packet Loss Prob.):\t%v\t|\n", simulator.results.lossProbability)
	fmt.Printf("| ρ (Utilization of queue): \t%v\t|", simulator.results.queueUtilization)
		fmt.Printf(" M (Simulation Iterations): \t%v\t|\n", M)
	fmt.Printf("| E[N] (Avg size of queue): \t%v\t|", simulator.results.avgQueueSize)
		fmt.Printf("\t\t\t\t\t|\n")
	fmt.Print("|-------------------------------------------------------------------------------|\n")
}
