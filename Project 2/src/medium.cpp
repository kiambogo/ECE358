#include "medium.hpp"
#include "signal.hpp"
#include "simulation.hpp"
#include <cassert>
#include <iostream>

void medium::propagate()
{
	for (auto it = signals.begin(); it != signals.end();) {
		(*it)->pos += (double)(*it)->dir * (double)propagation_delay * (double)sim->tick_length;
		assert((*it)->pos % 10 == 0);
		if ((*it)->pos < 0 || (*it)->pos > (int)((sim->n - 1) * sim->distance_between_nodes)) {
			delete *it;
			it = signals.erase(it);
		} else {
			++it;
		}
	}
}

void medium::add_signal(class signal *s)
{
	signals.push_back(s);
}

bool medium::signal_at_pos(unsigned int pos)
{
	for (unsigned int i = 0; i < signals.size(); i++) {
		if (signals[i]->pos == (int)pos) {
			return true;
		}
	}
	return false;
}
