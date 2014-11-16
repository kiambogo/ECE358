#pragma once

class medium;

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

  host(medium *network) : network(network) {}; 
  void run ();
  void transmit ();
  void sense ();
  void jam ();
};
