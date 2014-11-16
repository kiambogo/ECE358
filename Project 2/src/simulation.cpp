#include <list>
#include <map>
#include <cmath>
#include <random>
#include "simulation.hpp"
#include "host.hpp"
#include "medium.hpp"

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
	network = new medium(200000000, this);

	// Create hosts

	for (int i = 0; i < n; i++) {
		all_hosts.push_back(host());
	}

	// Fill generated_packets map
	unsigned int next_packet_generation_tick = 0;
	for (int i = 0; i < n-1; i++) {
		std::random_device rd;
		std::mt19937 gen(rd());
		std::uniform_real_distribution<> dis(0, 1);
		next_packet_generation_tick += ((-1/a) * log(dis(gen) - 1)) * (1/tick_length) + 1;
		while (next_packet_generation_tick < (run_time / tick_length)) {
			generated_packets.insert(std::pair<unsigned int,host>(next_packet_generation_tick, all_hosts[i]));
		}
	}
}

void simulation::tick()
{
	// Propagate signals
	network->propagate();

	// Check for generated packets / mark hosts active

	// Run host logic for each active host
}
