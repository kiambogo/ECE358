package main

// The Packet struct
type Packet struct {
	remainingBits int
	generatedAt uint32
}

func (p *Packet) decrementRemainingBits(rate int) {
  p.remainingBits -= rate
}
