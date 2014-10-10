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
	tickCounter       uint32
	timeToArrival     uint32
	results           SimulatorResults
}

type SimulatorResults struct {
	packetsReceived   uint32
	droppedPackets    uint32
	sentPackets       uint32
	idleTicks         uint32
	summedQueueSize   uint32
	summedSojurnTime  uint32
	avgSojurnTime     uint32
	avgQueueSize      uint32
	lossProbability   float32
	queueUtilization  float64
}

func (simulator *Simulator) initializeSimulator() {
	simulator.queue = Queue{list.New()} // Reset the queue
	simulator.tickCounter = 0
	simulator.timeToArrival = 0
	simulator.results.packetsReceived = 0
	simulator.results.summedQueueSize = 0
	simulator.results.summedSojurnTime = 0
	simulator.results.droppedPackets = 0
	simulator.results.sentPackets = 0
	simulator.results.idleTicks = 0
	simulator.results.lossProbability = 0

	simulator.results.avgSojurnTime = 0
	simulator.results.avgQueueSize = 0

	calculateArrival(simulator)
}

func (s *Simulator) startSimulation() {
	s.initializeSimulator()

	for (uint32(s.runTime) - s.tickCounter > 0) {
		packetArrival(s)
		packetDeparture(s)
		updateCalculations(s)

		s.tickCounter++
	}
}

func calculateArrival(s *Simulator) {
	s.timeToArrival = uint32(((-1/s.lambda) * math.Log(float64(1)-randGenerator()))*100000)
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

func (s *Simulator) computeResults() {
	s.results.avgQueueSize = s.results.summedQueueSize / uint32(s.runTime)
	s.results.avgSojurnTime = s.results.summedSojurnTime / s.results.packetsReceived
	s.results.lossProbability = float32(s.results.droppedPackets / (s.results.packetsReceived+s.results.droppedPackets))
	s.results.queueUtilization = float64(float64(s.packetSize) * (s.lambda/float64(s.transmissionRate)))
}
