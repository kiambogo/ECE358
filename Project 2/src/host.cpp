#include "host.hpp"
#include "signal.hpp"
#include "medium.hpp"
#include "simulation.hpp"
#include <cassert>
#include <iostream>

const unsigned int JAMMING_BITS = 48;

int host::run()
{
	int ret = 0;
	switch (state) {
	case WAIT: // Initial state for new packet arrival
		wait();
		break;
	case TRANSMIT: // State for transmitting
		ret = transmit();
		break;
	case JAM: // State for jamming
		jam();
		break;
	}
	return ret;
}

int host::transmit()
{
	int ret = 0;
	if (network->signal_at_pos(position)) {  // If there is another signal at this node
		std::cout << this << " Moving to JAM state\n";
		state = JAM;
		bit_time_counter = JAMMING_BITS * (1. / sim->w) * (1. / sim->tick_length);
	} else {   // Transmit
		network->add_signal(new signal(position, false, signal::RIGHT));
		network->add_signal(new signal(position, false, signal::LEFT));
		bit_time_counter--;
		if (bit_time_counter == 0) {
			num_packets--;
			if (num_packets == 0) {
				std::cout << this << " Moving to WAIT state\n";
				active = false;
				state = WAIT;
				ret = 1;
			} else {
				// Keep transmitting the next packet
				bit_time_counter = (sim->l * 8.) / sim->w * (1./sim->tick_length);
			}
		}
	}
	return ret;
}

void host::wait() {
	if (!network->signal_at_pos(position)) { // Channel is clear
		assert(state == WAIT);
		std::cout << this << " Moving to TRANSMIT state\n";
		state = TRANSMIT;
		bit_time_counter = ((sim->l * 8.) / sim->w) * (1./sim->tick_length);
		std::cout << bit_time_counter << "\n";
	}
}

void host::jam() {
	network->add_signal(new signal(position, true, signal::RIGHT));
	network->add_signal(new signal(position, true, signal::LEFT));
	bit_time_counter--;
	if (bit_time_counter == 0) {
		// TODO: Implement exponential backoff
		state = WAIT;
	}
}
