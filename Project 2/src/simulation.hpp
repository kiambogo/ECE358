#pragma once

#include <list>
#include <map>
#include "host.hpp"
#include "signal.hpp"

class simulation
{
public:
	const unsigned int n;
	const unsigned int a;
	static const unsigned int w = 1000000; // 1 Mbps
	static const unsigned int l = 1500;
	const unsigned int p;

	const static double tick_length = 50 * 10^-9; // seconds

	simulation(unsigned int n, unsigned int a, unsigned int p, unsigned int run_time) : n(n), a(a), p(p), run_time(run_time) {};
	void run();

private:
	unsigned int ticks;
	unsigned int run_time;

	medium *network;
	std::list<host> active_hosts;
	std::map<unsigned int, host> generated_packets;

	void init();
	void tick();
};

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

	// Fill generated_packets map
}

void simulation::tick()
{
	// Propagate signals

	// Check for generated packets / mark hosts active

	// Run host logic for each active host
}
