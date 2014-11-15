#pragma once

class signal
{
public:
	enum direction {LEFT = -1, RIGHT = 1};
	bool jamming;

	signal(int pos, bool jamming, direction dir, unsigned int n) : pos(pos), jamming(jamming), dir(dir), n(n) {};
	void propagate();

private:
	const direction dir;
	int pos;
	const int n;
};
