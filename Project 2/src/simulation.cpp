#include <list>
#include <map>
#include <cmath>
#include "simulation.hpp"
#include "host.hpp"

void simulation::run()
{
	init();
	for (ticks = 0; ticks < run_time / tick_length; ticks++) {
		tick();
	}
}

void simulation::init()
{
	// Create medium

	// Create hosts

	/**host all_hosts [n] = {}

	// Fill generated_packets map
	for (int i = 0; i < n-1; i++) {
		int next_packet = ((-1/a) * log(rand() % 1)) * (1/tick_length)
		while(next_packet < (1/tick_length)) {
			generated_packets.insert(next_packet, all_hosts[i])
		}
	}*/
}

void simulation::tick()
{
	// Propagate signals

	// Check for generated packets / mark hosts active

	// Run host logic for each active host
}
