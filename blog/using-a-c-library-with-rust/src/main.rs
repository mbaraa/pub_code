#[link(name = "alsa", kind = "static")]
extern "C" {
    fn init_alsa() -> i32;
    fn destroy_alsa() -> i32;
    fn play_frequency(freq: f32, rate: u16, latency: f32, duration: f32) -> i32;
}

enum NoteDuration {
    TwoNotes,
    WholeNote,
    HalfNote,
}

impl NoteDuration {
    fn value(&self, secs: f32) -> f32 {
        match *self {
            Self::TwoNotes => 2.0 * secs,
            Self::WholeNote => secs,
            Self::HalfNote => secs * 0.5,
        }
    }
}

struct Note {
    freq: f32,
    duration: NoteDuration,
}

impl Note {
    fn new(freq: f32, duration: NoteDuration) -> Self {
        Self { freq, duration }
    }
}

fn main() {
    unsafe {
        init_alsa();

        vec![
            Note::new(110.0, NoteDuration::HalfNote),
            Note::new(174.61, NoteDuration::WholeNote),
            Note::new(110.0, NoteDuration::HalfNote),
            Note::new(164.81, NoteDuration::WholeNote),
            Note::new(110.0, NoteDuration::HalfNote),
            Note::new(174.61, NoteDuration::HalfNote),
            Note::new(196.0, NoteDuration::WholeNote),
            Note::new(164.81, NoteDuration::HalfNote),
            Note::new(110.0, NoteDuration::WholeNote),
            Note::new(196.0, NoteDuration::WholeNote),
            Note::new(174.61, NoteDuration::WholeNote),
            Note::new(164.81, NoteDuration::HalfNote),
            Note::new(146.83, NoteDuration::WholeNote),
            Note::new(164.81, NoteDuration::HalfNote),
            Note::new(130.81, NoteDuration::HalfNote),
            Note::new(220.0, NoteDuration::WholeNote),
            Note::new(130.81, NoteDuration::HalfNote),
            Note::new(196.0, NoteDuration::WholeNote),
            Note::new(130.81, NoteDuration::HalfNote),
            Note::new(220.0, NoteDuration::WholeNote),
            Note::new(233.08, NoteDuration::WholeNote),
            Note::new(196.0, NoteDuration::WholeNote),
            Note::new(220.0, NoteDuration::HalfNote),
            Note::new(233.08, NoteDuration::WholeNote),
            Note::new(220.0, NoteDuration::WholeNote),
            Note::new(196.0, NoteDuration::WholeNote),
            Note::new(174.61, NoteDuration::WholeNote),
            Note::new(174.61, NoteDuration::TwoNotes),
            Note::new(164.81, NoteDuration::WholeNote),
        ]
        .iter()
        .for_each(|note| {
            println!("{}", note.freq);
            play_frequency(note.freq, 44100, 0.1, note.duration.value(1.0));
        });

        destroy_alsa();
    }
}
