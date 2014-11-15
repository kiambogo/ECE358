package main

// The Packet struct
type Packet struct {
	remainingBits uint32
	generatedAt uint64
}

func (p *Packet) decrementRemainingBits(rate uint32) {
  p.remainingBits -= rate
}
