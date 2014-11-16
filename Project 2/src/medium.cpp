#include "medium.hpp"
#include "signal.hpp"
#include "simulation.hpp"

void medium::propagate()
{
	// Old list-based implementation
	/*for (std::list<signal>::iterator it = signals.begin(); it != signals.end(); ++it) {
		it->pos += it->dir * propagation_delay * sim->tick_length;
		if (it->pos < 0 || it->pos > (sim->n - 1) * sim->distance_between_nodes) {
			signals.erase(it);
		}
	}*/
	for (int i = 0; i < signals.size(); i++) {
		signals[i].pos += signals[i].dir * propagation_delay * sim->tick_length;
		if (signals[i].pos < 0 || signals[i].pos > (sim->n - 1) * sim->distance_between_nodes) {
			signals.erase(signals.begin() + i);
		}
	}
}

void medium::add_signal(signal *s)
{
	signals.push_back(*s);
}

bool medium::signal_at_pos(unsigned int pos)
{
	for (int i = 0; i < signals.size(); i++) {
		if (signals[i].pos == pos) {
			return true;
		}
	}
	return false;
}
