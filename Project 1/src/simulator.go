package main

import (
	"math"
	"container/list"
)

type Simulator struct {
	lambda            float32
	transmissionRate  uint32
	packetSize        uint32
	bufferSize        int
	ticks							uint64
	tickDuration			uint64
	queue             Queue
	tickCounter       uint64
	timeToArrival     uint64
	results           SimulatorResults
}

type SimulatorResults struct {
	packetsReceived       uint32
	droppedPackets        uint32
	sentPackets           uint32
	idleTicks             uint64
	summedQueueSize       float64
	summedSojurnTime      float64
	avgSojurnTime         float64
	avgQueueSize          float64
	idleServerProportion  float64
	lossProbability       float64
	queueUtilization      float64
}

func (simulator *Simulator) initializeSimulator() {
	simulator.queue = Queue{list.New()} // Reset the queue
	simulator.tickCounter = 0
	simulator.timeToArrival = 0

	simulator.results.packetsReceived = 0
	simulator.results.droppedPackets = 0
	simulator.results.sentPackets = 0
	simulator.results.idleTicks = 0
	simulator.results.summedQueueSize = 0
	simulator.results.summedSojurnTime = 0
	simulator.results.avgSojurnTime = 0
	simulator.results.avgQueueSize = 0
	simulator.results.idleServerProportion = 0
	simulator.results.lossProbability = 0
	simulator.results.queueUtilization = 0

	calculateArrival(simulator)
}

func (s *Simulator) startSimulation() {
	s.initializeSimulator()

	for (uint64(s.ticks) - s.tickCounter-1 > 0) {
		packetArrival(s)
		packetDeparture(s)
		updateCalculations(s)

		s.tickCounter++
	}
}

func calculateArrival(s *Simulator) {
	s.timeToArrival = uint64((float64(-1/s.lambda) * math.Log(float64(1)-randGenerator()))*float64(s.tickDuration))
}

func packetArrival(s *Simulator) {
	if (s.timeToArrival > 0) {
		s.timeToArrival--
	} else {
		if (s.bufferSize == -1 || s.queue.buffer.Len() < s.bufferSize) {
			s.queue.enqueue(&Packet{remainingBits:s.packetSize, generatedAt:s.tickCounter})
			s.results.packetsReceived++
		} else {
			s.results.droppedPackets++
		}
		calculateArrival(s)
	}
}

func packetDeparture(s *Simulator) {
	if (s.queue.buffer.Len() == 0) {
		s.results.idleTicks++
	} else {
		s.queue.peek().decrementRemainingBits(s.transmissionRate)
		if (s.queue.peek().remainingBits <= 0) {
			s.results.sentPackets++
			s.results.summedSojurnTime += float64(uint32(s.tickCounter) - uint32(s.queue.peek().generatedAt))
			s.queue.dequeue()
		}
	}
}

func updateCalculations(s *Simulator) {
	s.results.summedQueueSize += float64(s.queue.buffer.Len())
}

func (s *Simulator) computeResults() {
	s.results.avgQueueSize = float64(s.results.summedQueueSize / float64(s.ticks))
	if (s.results.packetsReceived == 0) { s.results.avgSojurnTime = 0
  	} else { s.results.avgSojurnTime = s.results.summedSojurnTime / float64(s.results.packetsReceived) }
	if (s.results.packetsReceived+s.results.droppedPackets == 0) { s.results.lossProbability = 0
	} else { s.results.lossProbability = float64(s.results.droppedPackets) / float64(s.results.packetsReceived+s.results.droppedPackets)}
	s.results.queueUtilization = float64(float32(s.packetSize) * (s.lambda/float32(s.transmissionRate)))
	s.results.idleServerProportion = float64(float64(s.results.idleTicks) / float64(s.ticks))
}
