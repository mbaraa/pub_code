#include "alsa.h"

snd_pcm_t *handle;

long long second_to_micro(float seconds) { return (long long)(seconds * 1e6); }

int init_alsa() {
  return snd_pcm_open(&handle, "default", SND_PCM_STREAM_PLAYBACK,
                      0 /* blocked mode */);
}

int destroy_alsa() { return snd_pcm_close(handle); }

int play_frequency(float freq, u_int16_t rate, float latency, float duration) {
  if (latency > 150.0) {
    latency = 150.0;
  }
  latency = second_to_micro(latency);

  unsigned char buffer[(int)(rate * duration)];

  for (int i = 0; i < sizeof(buffer); i++) {
    buffer[i] = 0xFF * sin(2 * M_PI * freq * i / rate);
  }

  if (0 != snd_pcm_set_params(handle, SND_PCM_FORMAT_U8,
                              SND_PCM_ACCESS_RW_INTERLEAVED, 1 /* channels */,
                              rate /* rate [Hz] */, 1 /* soft resample */,
                              latency /* latency [us] */)) {
    return 0;
  }

  snd_pcm_writei(handle, buffer, sizeof(buffer));

  return 0;
}
