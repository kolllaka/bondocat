#include <stdio.h>
#include <stdlib.h>

#define _WIN32_WINNT_WIN10 0x0A00

#include <windows.h>

int fx()
{
	LPPOINT pPnt;
	pPnt = malloc(sizeof(*pPnt));

	GetCursorPos(pPnt);

	int posX = pPnt[0].x;

	free(pPnt);

	return posX;
}

int fy()
{
	LPPOINT pPnt;
	pPnt = malloc(sizeof(*pPnt));

	GetCursorPos(pPnt);

	int posY = pPnt[0].y;

	free(pPnt);

	return posY;
}

int f()
{
	// q
	if (GetKeyState(0x51) < 0)
		return 1;
	// a
	else if (GetKeyState(0x41) < 0)
		return 2;
	// z
	else if (GetKeyState(0x5A) < 0)
		return 3;
	// w
	else if (GetKeyState(0x57) < 0)
		return 4;
	// s
	else if (GetKeyState(0x53) < 0)
		return 5;
	// x
	else if (GetKeyState(0x58) < 0)
		return 6;
	// e
	else if (GetKeyState(0x45) < 0)
		return 7;
	// d
	else if (GetKeyState(0x44) < 0)
		return 8;
	// c
	else if (GetKeyState(0x43) < 0)
		return 9;
	// r
	else if (GetKeyState(0x52) < 0)
		return 10;
	// f
	else if (GetKeyState(0x46) < 0)
		return 11;
	// v
	else if (GetKeyState(0x56) < 0)
		return 12;
	// t
	else if (GetKeyState(0x54) < 0)
		return 13;
	// g
	else if (GetKeyState(0x47) < 0)
		return 14;
	// b
	else if (GetKeyState(0x42) < 0)
		return 15;
	// y
	else if (GetKeyState(0x59) < 0)
		return 16;
	// h
	else if (GetKeyState(0x48) < 0)
		return 17;
	// n
	else if (GetKeyState(0x4E) < 0)
		return 18;
	// u
	else if (GetKeyState(0x55) < 0)
		return 19;
	// j
	else if (GetKeyState(0x4A) < 0)
		return 20;
	// m
	else if (GetKeyState(0x4D) < 0)
		return 21;
	// i
	else if (GetKeyState(0x49) < 0)
		return 22;
	// k
	else if (GetKeyState(0x4B) < 0)
		return 23;
	// o
	else if (GetKeyState(0x53) < 0)
		return 24;
	// l
	else if (GetKeyState(0x53) < 0)
		return 25;
	// p
	else if (GetKeyState(0x53) < 0)
		return 26;
}