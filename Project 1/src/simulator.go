package main

import (
  "fmt"
	"math"
)

type Simulator struct {
  lambda float64
  transmissionRate int
  packetSize int
  runTime int
	bufferSize int

	timeToArrival int
	tickCounter int
	droppedPackets int
	idleTicks int
	queue Queue
}

func (simulator Simulator) initializeSimulator() {
	simulator.tickCounter = 0
	simulator.droppedPackets = 0
	simulator.idleTicks = 0
}

func (simulator Simulator) startSimulation() {
  fmt.Println("yo")
	simulator.initializeSimulator()

	for (simulator.runTime - simulator.tickCounter > 0) {

	}

  calculateArrival(&simulator)
}

func calculateArrival(s *Simulator) (int) {
	return int((-1/s.lambda) * math.Log(float64(1)-randGenerator()))
}

func packetArrival(s *Simulator) {
	if (s.timeToArrival == 0) {
		if (s.bufferSize == -1 || s.queue.buffer.Len() < s.bufferSize) {
			s.queue.enqueue(Packet{remainingBits:s.packetSize, generatedAt:s.tickCounter})
		} else {
			s.droppedPackets++
		}
		calculateArrival(s)
	} else {
		s.timeToArrival--
	}
}

func packetDeparture(s *Simulator) {
	if (s.queue.buffer.Len() == 0) {
		s.idleTicks++
	} else {
		s.queue.peek().decrementRemainingBits(s.transmissionRate)
	}
}
