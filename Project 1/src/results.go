package main

type Results struct {
	summedIdleTicks        uint32
	summedAvgQueueSize     uint32
	summedAvgSojournTime   uint32
	summedLossProbability  float32
	summedQueueUtilization float64

	finalIdleTicks        uint32
	finalAvgQueueSize     uint32
	finalAvgSojournTime   uint32
	finalLossProbability  float32
	finalQueueUtilization float64
}

func (r *Results) updateResultSet(s *Simulator) {
	r.summedIdleTicks += s.results.idleTicks
	r.summedAvgQueueSize += s.results.avgQueueSize
	r.summedAvgSojournTime += s.results.avgSojurnTime
	r.summedLossProbability += s.results.lossProbability
	r.summedQueueUtilization += s.results.queueUtilization
}

func (r *Results) computeFinalResults(m int) {
	r.finalIdleTicks = r.summedIdleTicks/uint32(m)
	r.finalAvgQueueSize = r.summedAvgQueueSize/uint32(m)
	r.finalAvgSojournTime = r.summedAvgSojournTime/uint32(m)
	r.finalLossProbability = r.summedLossProbability/float32(m)
	r.finalQueueUtilization = r.summedQueueUtilization/float64(m)
}
