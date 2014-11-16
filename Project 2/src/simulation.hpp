#pragma once

#include <vector>
#include <map>
#include <cmath>
#include "host.hpp"

class medium;
class host;

class simulation
{
public:
	const unsigned int n;
	const unsigned int a;
	const unsigned int w = 1000000; // 1 Mbps
	const unsigned int l = 1500;
	const unsigned int p;

	static const unsigned int distance_between_nodes = 10;

	const double tick_length = .00000005; // 50 nanoseconds

	unsigned int successful_packet_transmissions;
	unsigned int ticks;
	std::vector<unsigned int> packet_transmission_delays;

	unsigned int debug_wait_state_cnt;

	simulation(unsigned int n, unsigned int a, unsigned int p, double run_time) : n(n), a(a), p(p), successful_packet_transmissions(0), debug_wait_state_cnt(0), run_time(run_time) {};
	void run();

private:
	double run_time; // seconds

	medium *network;
	std::vector<host *> all_hosts;
	std::vector<host *> active_hosts;
	std::multimap<unsigned int, host *> generated_packets;

	void init();
	void tick();
};
