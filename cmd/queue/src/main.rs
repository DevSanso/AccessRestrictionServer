mod buffer;
mod adapter;
mod env;
mod timer;
mod kafka;

fn main() {
    let tim = timer::Timer::new(20).start();
    loop {

    }
}
