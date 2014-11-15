/*
  ECE 358 Project 2
  Christopher Poenaru  |  cpoenaru  |  20409287
  Paul Trautrim        |  pctrautr  |  20348861

  This code is original and is the work of us as partners.
*/

package main

// N: the number of computers connected to the LAN (variable)
// A: Data packets arrive at the MAC layer following a Poisson process with an average arrival rate of A packets/second (variable)
// W: the speed of the LAN (fixed)
// L: packet length (fixed)
// P: Persistence parameter for P-persistent CSMA protocols.

func main() {
	question1()
}

func question1() {
	ticks := uint64(60000000000)
	tickDuration := uint64(1000000000)
	N := 1
	A := 1
}
