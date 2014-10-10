package main

import (
  "fmt"
	"math"
)

type Simulator struct {
	lambda            float64
	transmissionRate  int
	packetSize        int
	runTime           int
	bufferSize        int
	queue             Queue
	results           SimulatorResults
	tickCounter       int
	timeToArrival     int
}

type SimulatorResults struct {
	droppedPackets  int
	sentPackets     int
	idleTicks       int
	summedQueueSize uint32

	avgQueueSize uint32
}

func (simulator *Simulator) initializeSimulator() {
	simulator.tickCounter = 0
	simulator.results.droppedPackets = 0
	simulator.results.sentPackets = 0
	simulator.results.idleTicks = 0

	calculateArrival(simulator)
}

func (simulator *Simulator) startSimulation() {
	fmt.Printf("Starting Simulator with λ=%v, L=%v, C=%v, bufferSize=%v ...\n",
		simulator.lambda, simulator.packetSize, simulator.transmissionRate, simulator.bufferSize)
	simulator.initializeSimulator()

	for (simulator.runTime - simulator.tickCounter > 0) {
		packetArrival(simulator)
		packetDeparture(simulator)
		updateCalculations(simulator)

		simulator.tickCounter++
	}
	fmt.Printf("Completed Simulation\n")
}

func calculateArrival(s *Simulator) {
	s.timeToArrival = int(((-1/s.lambda) * math.Log(float64(1)-randGenerator()))*100000)
}

func packetArrival(s *Simulator) {
	if (s.timeToArrival == 0) {
		if (s.bufferSize == -1 || s.queue.buffer.Len() < s.bufferSize) {
			s.queue.enqueue(Packet{remainingBits:s.packetSize, generatedAt:s.tickCounter})
		} else {
			s.results.droppedPackets++
		}
		calculateArrival(s)
	} else {
		s.timeToArrival--
	}
}

func packetDeparture(s *Simulator) {
	if (s.queue.buffer.Len() == 0) {
		s.results.idleTicks++
	} else {
		s.queue.peek().decrementRemainingBits(s.transmissionRate)
		if (s.queue.peek().remainingBits <= 0) {
			s.results.sentPackets++
			s.queue.dequeue()
		}
	}
}

func updateCalculations(s *Simulator) {
	s.results.summedQueueSize += uint32(s.queue.buffer.Len())
}

func (simulator *Simulator) computeResults() {
	simulator.results.avgQueueSize = simulator.results.summedQueueSize/ uint32(simulator.runTime)
}
