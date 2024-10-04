#ifndef MEME
#define MEME
#include <stdlib.h>

typedef struct c_meme c_meme;

#ifdef __cplusplus
extern "C" {
#endif

// here we have some c methods that go can call to interface
// idea for derate would be to "create an ESS asset"
// then call the member methods
extern c_meme *createMeme(float data);
extern void destroyMeme(void *self);
extern void incrementMeme(c_meme *self);
extern void decrementMeme(c_meme *self);
extern float getMemeData(c_meme *self);

#ifdef __cplusplus
}
#endif

#endif
