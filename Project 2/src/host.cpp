#include "host.hpp"
#include "signal.hpp"
#include "medium.hpp"
#include "simulation.hpp"
#include <cassert>
#include <random>
#include <iostream>

host::host(simulation *sim, medium *network, unsigned int position) : active(false), position(position), sim(sim), network(network), state(SENSE), i(0), has_deferred(false)
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
		network->add_signal(new class signal(position, false, signal::RIGHT));
		network->add_signal(new class signal(position, false, signal::LEFT));
		bit_time_counter--;
		if (bit_time_counter == 0) {
			sim->successful_packet_transmissions++;
			i = 0;
			has_deferred = false;
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
			std::random_device rd;
			std::mt19937 gen(rd());
			std::uniform_real_distribution<> dis(0, 1);

			if (sim->p > 0 && dis(gen) >= sim->p) {
				// Defer
				has_deferred = true;
				state = WAIT;
				bit_time_counter = 2. * (double)(sim->n - 1) / (double)network->propagation_delay + 0.5;
			} else {
				// Transmit
				std::cout << sim->ticks << " " << this << " Moving to TRANSMIT state\n";
				state = TRANSMIT;
				bit_time_counter = ((sim->l * 8.) / sim->w) * (1./sim->tick_length);
				std::cout << sim->ticks << " " << bit_time_counter << "\n";
			}
		}
	} else {
		if (sim->p == 0) { // Non-persistent is a special case
			// Our random wait is the same as if a collision were detected
			bit_time_counter = calculate_random_backoff();
			state = WAIT;
		} else if (sim->p == 1) { // 1-persistent is a special case
			// Restart sensing time
			bit_time_counter = SENSING_BITS * (1. / sim->w) * (1. / sim->tick_length);
		} else {
			if (has_deferred) {
				state = WAIT;
				bit_time_counter = calculate_random_backoff();
				has_deferred = false;
			} else {
				// Wait until next slot
				state = WAIT;
				bit_time_counter = 2. * (double)(sim->n - 1) / (double)network->propagation_delay + 0.5;
			}
		}
	}
}

unsigned int host::calculate_random_backoff()
{
	if (i < KMAX) {
		i++;
	}
	std::random_device rd;
	std::mt19937 gen(rd());
	std::uniform_real_distribution<> dis(0, std::pow(2., (double)i) - 1);

	unsigned int r = dis(gen) + 0.5;
	unsigned int tb = (double)r * (double)TP * (1. / sim->w) * (1. / sim->tick_length) + 0.5;

	return tb;
}

void host::jam() {
	network->add_signal(new class signal(position, true, signal::RIGHT));
	network->add_signal(new class signal(position, true, signal::LEFT));
	bit_time_counter--;
	if (bit_time_counter == 0) {
		bit_time_counter = calculate_random_backoff();

		std::cout << sim->ticks << " " << this << " JAM finished, moving to WAIT state with counter " << bit_time_counter << "\n";
		sim->debug_wait_state_cnt++;
		std::cout << sim->ticks << " " << this << " Wait state cnt: " << sim->debug_wait_state_cnt << "\n";
		state = WAIT;
	}
}

void host::wait() {
	if (bit_time_counter == 0) {
		//std::cout << sim->ticks << " " << this << " WAIT finished, moving to SENSE state\n";
		bit_time_counter = SENSING_BITS * (1. / sim->w) * (1. / sim->tick_length);
		sim->debug_wait_state_cnt--;
		//std::cout << sim->ticks << " " << this << " Wait state cnt: " << sim->debug_wait_state_cnt << "\n";
		state = SENSE;
	}
	bit_time_counter--;
}
