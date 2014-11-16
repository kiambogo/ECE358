#include "host.hpp"
#include "signal.hpp"
#include "medium.hpp"
#include "simulation.hpp"

void host::run () {
  switch (state) {
    case 0: // Initial state for new packet arrival
      sense();
    break;
    case 1: // State for transmitting
      transmit();
    break;
    case 2: // State for jamming
      jam();
    break;
  }
}

void host::transmit() {
  if (num_packets > 0) {  // Sense
    state = 2;
    bit_time_counter = 48 * (1/sim->w);
  } else {   // Transmit
    network->add_signal(new signal(position, false, signal::RIGHT));
    network->add_signal(new signal(position, false, signal::LEFT));
  }
}

void host::sense() {
  if (retrieveSignals() == 0) {
    if(state == 0) {
      state = 1;
      bit_time_counter = ((sim->l * 8) / sim->w) * (1/sim->tick_length);
    }
  } else {
    if(state == 1) {
      state = 2;
      bit_time_counter = 48 * (1/sim->w) * 200;
    }
  }
}

void host::jam() {
  network->add_signal(new signal(position, true, signal::RIGHT));
  network->add_signal(new signal(position, true, signal::LEFT));
}
