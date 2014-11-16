#pragma once

class signal
{
public:
	int pos;
	enum direction {LEFT = -1, RIGHT = 1};
	bool jamming;
	direction dir;

	signal(int pos, bool jamming, direction dir) : pos(pos), jamming(jamming), dir(dir) {};
};
