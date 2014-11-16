#pragma once

class medium;
class simulation;

class host
{
public:
  int   num_packets;
  int   state;
  int   position;
  bool  active;
  long  transmission_counter;
  long  bit_time_counter;

  host(medium *network) : network(network) {};
  void run ();
  void transmit ();
  void sense ();
  void jam ();

private:
  medium *network;
  simulation *sim;
};
