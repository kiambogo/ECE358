package main

// The Packet struct
type Packet struct {
	remainingBits int
	generatedAt int
}

func (p *Packet) decrementRemainingBits(rate int) {
  p.remainingBits -= rate
}
