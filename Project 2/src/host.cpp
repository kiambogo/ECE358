#include "host.hpp"
#include "signal.hpp"
#include "medium.hpp"
#include "simulation.hpp"
#include <cassert>
#include <random>
#include <iostream>

#define ONE_PERSISTENT

host::host(simulation *sim, medium *network, unsigned int position) : sim(sim), network(network), state(SENSE), active(false), position(position), i(0)
{
	bit_time_counter = SENSING_BITS * (1. / sim->w) * (1. / sim->tick_length);
}

int host::run()
{
	int ret = 0;
	switch (state) {
	case SENSE: // Initial state for new packet arrival
		sense();
		break;
	case TRANSMIT: // State for transmitting
		ret = transmit();
		break;
	case JAM: // State for jamming
		jam();
		break;
	case WAIT: // State for exponential backoff
		wait();
		break;
	}
	return ret;
}

int host::transmit()
{
	int ret = 0;
	if (network->signal_at_pos(position)) {  // If there is another signal at this node
		std::cout << sim->ticks << " " << this << " Moving to JAM state\n";
		state = JAM;
		bit_time_counter = JAMMING_BITS * (1. / sim->w) * (1. / sim->tick_length);
	} else {   // Transmit
		network->add_signal(new signal(position, false, signal::RIGHT));
		network->add_signal(new signal(position, false, signal::LEFT));
		bit_time_counter--;
		if (bit_time_counter == 0) {
			sim->successful_packet_transmissions++;
			i = 0;
			unsigned int arrival_tick = packet_arrival_times.front();
			sim->packet_transmission_delays.push_back(sim->ticks - arrival_tick);
			packet_arrival_times.erase(packet_arrival_times.begin());
			if (packet_arrival_times.empty()) {
				active = false;
				ret = 1;
			}
			state = SENSE;
			bit_time_counter = SENSING_BITS * (1. / sim->w) * (1. / sim->tick_length);
		}
	}
	return ret;
}

void host::sense() {
	if (!network->signal_at_pos(position)) { // Channel is clear
		bit_time_counter--;
		if (bit_time_counter == 0) {
			std::cout << sim->ticks << " " << this << " Moving to TRANSMIT state\n";
			state = TRANSMIT;
			bit_time_counter = ((sim->l * 8.) / sim->w) * (1./sim->tick_length);
			std::cout << sim->ticks << " " << bit_time_counter << "\n";
		}
	} else {
#ifdef ONE_PERSISTENT
		// Restart sensing time
		bit_time_counter = SENSING_BITS * (1. / sim->w) * (1. / sim->tick_length);
#elif NON_PERSISTENT
		std::random_device rd;
		std::mt19937 gen(rd());
		std::uniform_real_distribution<> dis(0, std::pow(2., (double)i) - 1);
		bit_time_counter = SENSING_BITS * (1. / sim->w) * (1. / sim->tick_length);
		state = WAIT;
#else
		assert(false);
#endif
	}
}

void host::jam() {
	network->add_signal(new signal(position, true, signal::RIGHT));
	network->add_signal(new signal(position, true, signal::LEFT));
	bit_time_counter--;
	if (bit_time_counter == 0) {
		if (i < KMAX) {
			i++;
		}
		std::random_device rd;
		std::mt19937 gen(rd());
		std::uniform_real_distribution<> dis(0, std::pow(2., (double)i) - 1);

		unsigned int r = dis(gen) + 0.5;
		unsigned int tb = (double)r * (double)TP * (1. / sim->w) * (1. / sim->tick_length) + 0.5;
		std::cout << this << " r: " << r << " tb: " << tb << "\n";
		bit_time_counter = tb;

		std::cout << sim->ticks << " " << this << " JAM finished, moving to WAIT state with counter " << bit_time_counter << "\n";
		sim->debug_wait_state_cnt++;
		std::cout << sim->ticks << " " << this << " Wait state cnt: " << sim->debug_wait_state_cnt << "\n";
		state = WAIT;
	}
}

void host::wait() {
	if (bit_time_counter == 0) {
		std::cout << sim->ticks << " " << this << " WAIT finished, moving to SENSE state\n";
		bit_time_counter = SENSING_BITS * (1. / sim->w) * (1. / sim->tick_length);
		sim->debug_wait_state_cnt--;
		std::cout << sim->ticks << " " << this << " Wait state cnt: " << sim->debug_wait_state_cnt << "\n";
		state = SENSE;
	}
	bit_time_counter--;
}
