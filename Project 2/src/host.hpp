#pragma once

#include "signal.hpp"
#include "medium.hpp"

class host
{
public:
  int   num_packets;
  int   state;
  int   position;
  bool  active;
  long  transmission_counter;
  long  bit_time_counter;
  medium *network;

  host(*network medium)
  void run ();
  void transmit ();
  void sense ();
  void jam ();
};

void host::run () {
  switch (state)
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
