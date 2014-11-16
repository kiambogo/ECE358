#pragma once

class signal
{
public:
	enum direction {LEFT = -1, RIGHT = 1};
	bool jamming;
	direction dir;
	int pos;

	signal(int pos, bool jamming, direction dir) : pos(pos), jamming(jamming), dir(dir) {};
};
