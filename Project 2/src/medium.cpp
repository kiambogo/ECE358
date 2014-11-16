#include <list>
#include "signal.hpp"
#include "simulation.hpp"

void medium::propagate()
{
	for (std::list<signal>::iterator it = signals.begin(); it != signals.end(); ++it) {
		it->pos += it->dir * propagation_delay * sim->tick_length;
		if (it->pos < 0 || it->pos > (sim->n - 1) * sim->distance_between_nodes) {
			signals.erase(it);
			delete it;
		}
	}
}

void medium::add_signal(signal *s)
{
	assert (s != NULL);
	signals.push_back(s);
}
