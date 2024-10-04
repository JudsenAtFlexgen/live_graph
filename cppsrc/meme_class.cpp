#include "meme_class.h"

meme::meme() { data = 0; }
meme::meme(float data) : data(data) {}
void meme::increment() { data++; }
void meme::decrement() { data--; }
