package main

type Results struct {
	summedReceivedPackets  uint32
	summedDroppedPackets   uint32
	summedSentPackets      uint32
	summedIdleTicks        uint32
	summedAvgQueueSize     uint32
	summedAvgSojournTime   uint32
	summedLossProbability  float64
	summedQueueUtilization float64

	finalReceivedPackets  uint32
	finalDroppedPackets   uint32
	finalSentPackets      uint32
	finalIdleTicks        uint32
	finalAvgQueueSize     uint32
	finalAvgSojournTime   uint32
	finalLossProbability  float64
	finalQueueUtilization float64
}

func (r *Results) updateResultSet(s *Simulator) {
	r.summedReceivedPackets += s.results.packetsReceived
	r.summedDroppedPackets += s.results.droppedPackets
	r.summedSentPackets += s.results.sentPackets
	r.summedIdleTicks += s.results.idleTicks
	r.summedAvgQueueSize += s.results.avgQueueSize
	r.summedAvgSojournTime += s.results.avgSojurnTime
	r.summedLossProbability += s.results.lossProbability
	r.summedQueueUtilization += s.results.queueUtilization
}

func (r *Results) computeFinalResults(m int) {
	r.finalReceivedPackets = r.summedReceivedPackets/uint32(m)
	r.finalDroppedPackets = r.summedDroppedPackets/uint32(m)
	r.finalSentPackets = r.summedSentPackets/uint32(m)
	r.finalIdleTicks = r.summedIdleTicks/uint32(m)
	r.finalAvgQueueSize = r.summedAvgQueueSize/uint32(m)
	r.finalAvgSojournTime = r.summedAvgSojournTime/uint32(m)
	r.finalLossProbability = r.summedLossProbability/float64(m)
	r.finalQueueUtilization = r.summedQueueUtilization/float64(m)
}
