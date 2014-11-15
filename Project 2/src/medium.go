package main

type Medium struct {
	nodes              map[int]Node
	carrierThreshold   int
	collisionThreshold int
	jammingSignal      int
	propagationSpeed   float64
}

func (m *Medium) updateVoltage(host string, value int) {
}

func (m *Medium) resetVoltage(host string) {
}

type MediumNode struct {
	prevSignals    []Signal
	currentSignals []Signal
	totalVoltage   int
}

type Signal struct {
	voltage  uint32
	prevNode *MediumNode
}
