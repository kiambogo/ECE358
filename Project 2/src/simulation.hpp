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

	const double tick_length = 50 * 10^-9; // seconds

	simulation(unsigned int n, unsigned int a, unsigned int p, unsigned int run_time) : n(n), a(a), p(p), run_time(run_time) {};
	void run();

private:
	unsigned int ticks;
	unsigned int run_time;

	medium *network;
	std::vector<host> all_hosts;
	std::vector<host> active_hosts;
	std::multimap<unsigned int, host> generated_packets;

	void init();
	void tick();
};
