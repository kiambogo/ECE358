package main

import (
	"math"
	"container/list"
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
	packetsReceived	 uint32
	droppedPackets   uint32
	sentPackets      int
	idleTicks        int
	summedQueueSize  uint32
	summedSojurnTime uint32
	avgSojurnTime    uint32
	avgQueueSize     uint32
	lossProbability	 float32
	queueUtilization float64
}

func (simulator *Simulator) initializeSimulator() {
	simulator.tickCounter = 0
	simulator.timeToArrival = 0
	simulator.results.packetsReceived = 0
	simulator.results.summedQueueSize = 0
	simulator.results.summedSojurnTime = 0
	simulator.results.droppedPackets = 0
	simulator.results.sentPackets = 0
	simulator.results.idleTicks = 0
	simulator.results.lossProbability = 0
	simulator.queue = Queue{list.New()}

	simulator.results.avgSojurnTime = 0
	simulator.results.avgQueueSize = 0

	calculateArrival(simulator)
}

func (simulator *Simulator) startSimulation() {
//	fmt.Printf("Starting Simulator with λ=%v, L=%v, C=%v, bufferSize=%v ...\n",
//		simulator.lambda, simulator.packetSize, simulator.transmissionRate, simulator.bufferSize)
	simulator.initializeSimulator()

	for (simulator.runTime - simulator.tickCounter > 0) {
		packetArrival(simulator)
		packetDeparture(simulator)
		updateCalculations(simulator)

		simulator.tickCounter++
	}
}

func calculateArrival(s *Simulator) {
	s.timeToArrival = int(((-1/s.lambda) * math.Log(float64(1)-randGenerator()))*100000)
}

func packetArrival(s *Simulator) {
	if (s.timeToArrival == 0) {
		if (s.bufferSize == -1 || s.queue.buffer.Len() < s.bufferSize) {
			s.queue.enqueue(&Packet{remainingBits:s.packetSize, generatedAt:s.tickCounter})
			s.results.packetsReceived++
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
			s.results.summedSojurnTime += (uint32(s.tickCounter) - uint32(s.queue.peek().generatedAt))
			s.queue.dequeue()
		}
	}
}

func updateCalculations(s *Simulator) {
	s.results.summedQueueSize += uint32(s.queue.buffer.Len())
}

func (simulator *Simulator) computeResults() {
	simulator.results.avgQueueSize = simulator.results.summedQueueSize / uint32(simulator.runTime)
	simulator.results.avgSojurnTime = simulator.results.summedSojurnTime / simulator.results.packetsReceived
	simulator.results.lossProbability = float32(simulator.results.droppedPackets / (simulator.results.packetsReceived+simulator.results.droppedPackets))
	simulator.results.queueUtilization = float64(float64(simulator.packetSize) * (simulator.lambda/float64(simulator.transmissionRate)))
}
