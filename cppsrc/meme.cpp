#include "meme.h"
#include "meme_class.h"
#include <cassert>
#include <stdlib.h>

struct c_meme {
  meme *spider_point;
};

extern c_meme *createMeme(float data) {
  c_meme *spider_point = (c_meme *)malloc(sizeof(c_meme));
  if (!spider_point) {
    return NULL;
  }

  spider_point->spider_point = new meme(data);
  return spider_point;
}

extern void destroyMeme(void *self) {
  assert(self);
  meme *spider_point = ((c_meme *)self)->spider_point;
  delete spider_point;
  free(self);
}

extern void incrementMeme(c_meme *self) {
  self->spider_point->increment();
}

extern void decrementMeme(c_meme *self) {
  self->spider_point->decrement();
}

extern float getMemeData(c_meme *self) {
  return self->spider_point->data;
}
