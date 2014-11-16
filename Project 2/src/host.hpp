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
	static const unsigned int TP = 512;
	static const unsigned int KMAX = 10;
	medium *network;
	simulation *sim;
	enum STATE {SENSE, TRANSMIT, JAM, WAIT};
	STATE state;
	unsigned int transmission_counter;
	unsigned int bit_time_counter;
	unsigned int i;

	int transmit();
	void sense();
	void jam();
	void wait();
};
