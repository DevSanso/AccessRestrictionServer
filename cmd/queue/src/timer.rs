use std::thread::{spawn, JoinHandle, sleep};
use std::sync::mpsc::{Sender, Receiver, channel, RecvError};
use std::sync::Once;
use std::time::Duration;


pub struct Timer {
    fps: u8,
    receiver: Option<Receiver<()>>,
}

pub struct TimerReceiver {
    recv : Receiver<()>
}

impl TimerReceiver {
    pub fn signal(&mut self) -> Result<(), RecvError> {self.recv.recv()}
}
impl Drop for  TimerReceiver {
    fn drop(mut self) { drop(self.recv); }
}


impl Timer {
    pub fn new(fps : u8) -> Timer { Timer {fps,receiver : None} }
    #[inline]
    fn thread_entry(fps : u8,sender : Sender<()>) -> Option<JoinHandle<()>> {
        let handle = spawn(move || {
            while sender.send(()).is_ok() {
                sleep(Duration::from_millis(fps as u64/1000));
            }
        });
        Some(handle)
    }
    pub fn start(mut self) -> TimerReceiver {
        let (send,recv) = channel::<()>();
        Timer::thread_entry(self.fps,send);
        TimerReceiver{recv}
    }

}

