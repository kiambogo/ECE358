#pragma once

class medium;
class simulation;

class host
{
public:
	int   num_packets;
	int   position;
	bool  active;

	host(simulation *sim, medium *network, unsigned int position) : sim(sim), network(network), state(WAIT), active(false), position(position), num_packets(0), transmission_counter(0), bit_time_counter(0) {};
	int run();

private:
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
