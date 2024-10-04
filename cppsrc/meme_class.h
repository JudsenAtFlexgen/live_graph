#ifndef MEMECLASS
#define MEMECLASS

// here we have some cpp type (class struct whatever)
#ifdef __cplusplus
struct meme {
  float data;

  meme();
  meme(float data);
  void increment();
  void decrement();
};
#endif
#endif
