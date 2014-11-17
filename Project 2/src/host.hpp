#pragma once

#include <vector>

class medium;
class simulation;

class host
{
public:
	std::vector<unsigned int> packet_arrival_times; // in ticks
	bool  active;
	int   position;

	host(simulation *sim, medium *network, unsigned int position);
	int run();

private:
	static const unsigned int JAMMING_BITS = 48;
	static const unsigned int SENSING_BITS = 96;
	static const unsigned int SLOT_BITS = 512;
	static const unsigned int TP = 512;
	static const unsigned int KMAX = 10;
	simulation *sim;
	medium *network;
	enum STATE {SENSE, TRANSMIT, JAM, WAIT};
	STATE state;
	unsigned int bit_time_counter;
	unsigned int i;
	bool has_deferred;

	int transmit();
	void sense();
	void jam();
	void wait();
	unsigned int calculate_random_backoff();
};
