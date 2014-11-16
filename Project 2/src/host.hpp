#pragma once

class medium;
class simulation;

class host
{
public:
	int   num_packets;
	int   position;
	bool  active;

	host(simulation *sim, medium *network, unsigned int position);
	int run();

private:
	static const unsigned int JAMMING_BITS = 48;
	static const unsigned int SENSING_BITS = 96;
	medium *network;
	simulation *sim;
	enum STATE {WAIT, TRANSMIT, JAM};
	STATE state;
	long transmission_counter;
	long bit_time_counter;

	int transmit();
	void wait();
	void jam();
};
