#pragma once

#include <vector>
#include "signal.hpp"
#include "simulation.hpp"

class simulation;

class medium
{
public:
	medium(unsigned int propagation_delay, simulation *sim) : propagation_delay(propagation_delay), sim(sim) {};
	void propagate();
	void add_signal(class signal *s);
	bool signal_at_pos(unsigned int pos);

private:
	const unsigned int propagation_delay;
	std::vector<class signal> signals;
	simulation *sim;
};
