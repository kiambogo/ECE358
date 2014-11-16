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

	std::cout << "End of simulation. Successful packet transmissions: " << successful_packet_transmissions << " Throughput: " << (successful_packet_transmissions * l * 8) / ((double)w * (double)run_time) << "\n";

	long double avg_packet_transmission_delay = 0;
	for (unsigned int i = 0; i < packet_transmission_delays.size(); i++) {
		avg_packet_transmission_delay += packet_transmission_delays[i] * tick_length;
	}
	avg_packet_transmission_delay /= packet_transmission_delays.size();
	std::cout << "Avg packet transmission delay: " << avg_packet_transmission_delay << " s\n";

	for (unsigned int i = 0; i < all_hosts.size(); i++) {
		delete all_hosts[i];
	}
	delete network;
}

void simulation::init()
{
	// Create medium
	network = new medium(200000000, this);

	// Create hosts
	for (unsigned int i = 0; i < n; i++) {
		all_hosts.push_back(new host(this, network, i * distance_between_nodes));
	}

	// Fill generated_packets map
	for (unsigned int i = 0; i < n; i++) {
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
		it->second->packet_arrival_times.push_back(ticks);
		if (it->second->active == false) {
			std::cout << "Moving " << it->second << " to active\n";
			it->second->active = true;
			active_hosts.push_back(it->second);
			std::cout << "Active hosts: " << active_hosts.size() << "\n";
		}
	}

	// Run host logic for each active host
	for (std::vector<host *>::iterator it = active_hosts.begin(); it != active_hosts.end();) {
		int ret = (*it)->run();
		if (ret) { // Remove node from inactive list
			std::cout << *it << " Done\n";
			it = active_hosts.erase(it);
			std::cout << "Active hosts: " << active_hosts.size() << "\n";
		} else {
			++it;
		}
	}
}
