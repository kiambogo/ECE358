package main

type Results struct {
	summedReceivedPackets  uint32
	summedDroppedPackets   uint32
	summedSentPackets      uint32
	summedIdleTicks        uint64
	summedAvgQueueSize     float64
	summedAvgSojournTime   float64
	summedIdleServerProp	 float64
	summedLossProbability  float64
	summedQueueUtilization float64

	finalReceivedPackets  uint32
	finalDroppedPackets   uint32
	finalSentPackets      uint32
	finalIdleTicks        uint64
	finalAvgQueueSize     float64
	finalAvgSojournTime   float64
	finalIdleServerProp	  float64
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
	r.summedIdleServerProp += s.results.idleServerProportion
	r.summedLossProbability += s.results.lossProbability
	r.summedQueueUtilization += s.results.queueUtilization
}

func (r *Results) computeFinalResults(m int) {
	r.finalReceivedPackets = r.summedReceivedPackets/uint32(m)
	r.finalDroppedPackets = r.summedDroppedPackets/uint32(m)
	r.finalSentPackets = r.summedSentPackets/uint32(m)
	r.finalIdleTicks = r.summedIdleTicks/uint64(m)
	r.finalAvgQueueSize = r.summedAvgQueueSize/float64(m)
	r.finalAvgSojournTime = r.summedAvgSojournTime/float64(m)
	r.finalIdleServerProp = r.summedIdleServerProp/float64(m)
	r.finalLossProbability = r.summedLossProbability/float64(m)
	r.finalQueueUtilization = r.summedQueueUtilization/float64(m)
}
