#pragma once

#include <list>
#include "signal.hpp"
#include "simulation.hpp"

class medium
{
public:
	medium(unsigned int propagation_delay, simulation *sim) : propagation_delay(propagation_delay), sim(sim) {};
	void propagate();

private:
	const unsigned int propagation_delay;
	std::list<signal> signals;
	simulation *sim;
};
