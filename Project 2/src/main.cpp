#include "simulation.hpp"

void question1();

int main()
{
	question1();

	return 0;
}

void question1()
{
	// N=20, A=5, P=1, time=60s
	simulation *sim = new simulation(20, 5, 0.5, 10);
	sim->run();
	delete sim;
	// N=40, A=5, P=1, time=60s
	// N=60, A=5, P=1, time=60s
	// N=80, A=5, P=1, time=60s
	// N=100, A=5, P=1, time=60s

	// N=20, A=6, P=1, time=60s
	// N=40, A=6, P=1, time=60s
	// N=60, A=6, P=1, time=60s
	// N=80, A=6, P=1, time=60s
	// N=100, A=6, P=1, time=60s

	// N=20, A=7, P=1, time=60s
	// N=40, A=7, P=1, time=60s
	// N=60, A=7, P=1, time=60s
	// N=80, A=7, P=1, time=60s
	// N=100, A=7, P=1, time=60s
}
