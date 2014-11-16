#include <map>
#include <cmath>
#include <random>
#include "simulation.hpp"
#include "host.hpp"
#include "medium.hpp"
#include <iostream>

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
		all_hosts.push_back(new host(this, network, i * distance_between_nodes));
	}

	// Fill generated_packets map
	for (int i = 0; i < n-1; i++) {
		std::random_device rd;
		std::mt19937 gen(rd());
		std::uniform_real_distribution<> dis(0, 1);

		// Generate the time until a new packet arrives, in ticks
		unsigned int next_packet_generation_tick = ((-1. / a) * log(1 - dis(gen))) * (1 / tick_length) + 0.5;

		// Do this for the entire simulation run time and map each arrival time to the node
		while (next_packet_generation_tick < (run_time / tick_length)) {
			generated_packets.insert(std::pair<unsigned int,host*>(next_packet_generation_tick, all_hosts[i]));
			next_packet_generation_tick += ((-1. / a) * log(1 - dis(gen))) * (1 / tick_length) + 0.5;
		}
	}
}

void simulation::tick()
{
	// Propagate signals
	network->propagate();

	// Check for generated packets / mark hosts active
	std::pair <std::multimap<unsigned int,host*>::iterator, std::multimap<unsigned int,host*>::iterator> ret;
	ret = generated_packets.equal_range(ticks);
	for (std::multimap<unsigned int,host*>::iterator it = ret.first; it != ret.second; ++it) {
		std::cout << "Packet generated at tick " << ticks << "\n";
		it->second->num_packets++;
		if (it->second->active == false) {
			std::cout << "Moving to active\n";
			it->second->active = true;
			active_hosts.push_back(it->second);
		}
	}

	// Run host logic for each active host
	for (std::vector<host *>::iterator it = active_hosts.begin(); it != active_hosts.end();) {
		int ret = (*it)->run();
		if (ret) { // Remove node from inactive list
			std::cout << "Done\n";
			it = active_hosts.erase(it);
		} else {
			++it;
		}
	}
}
