package main

// The Packet struct
type Packet struct {
	remainingBits int
	generatedAt int
}

func (p Packet) getRemainingBits() (int) {
	return p.remainingBits
}

func (p Packet) getGeneratedAt() (int) {
	return p.generatedAt
}

func (p Packet) decrementRemainingBits(rate int) {
  p.remainingBits -= rate
}
