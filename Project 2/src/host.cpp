#include "host.hpp"
#include "signal.hpp"
#include "medium.hpp"

void host::run () {
  switch (state) {
    case 0: // Initial state for new packet arrival
      sense();
    break;
    case 1: // State for transmitting
      transmit();
    break;
    case 2: // State for jamming
      jam()
    break;
  }
}

void host::transmit() {
  if (num_packets > 0) {  // Sense
    state = 2;
    bit_time_counter = 48 * BIT_TIME;
  } else {   // Transmit
    new signal(position, false, RIGHT, n);
    new signal(position, false, LEFT, n);
  }
}

void host::sense() {
    if (retrieveSignals() == 0) {
      if(state == 0) {
        state = 1;
        bit_time_counter = (((simulation -> l) * 8) / (simulation -> w)) * (1/(simulation -> tick_length));
      }
    } else {
      if(state == 1) {
        state = 2;
        bit_time_counter = 48 * (1/(simulation -> w)) * 200
      }
    }
}

void host::jam() {
  new signal(position, true, RIGHT, n);
  new signal(position, true, LEFT, n);
}