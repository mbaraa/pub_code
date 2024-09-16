#[link(name = "cairo", kind = "static")]
extern "C" {
    // the rusty protypes of the C functions.
    fn sq(num: f64);
}

pub fn square(num: f64) {
    unsafe {
        sq(num);
    }
}
