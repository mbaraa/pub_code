#ifndef ALSA_RS_H
#define ALSA_RS_H

#include <alsa/asoundlib.h>
#include <math.h>
#include <sys/types.h>

extern int init_alsa();
extern int destroy_alsa();
extern int play_frequency(float freq, u_int16_t rate, float latency,
                          float duration);

#endif
