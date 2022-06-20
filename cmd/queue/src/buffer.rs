use std::io::{Result, Error, ErrorKind::OutOfMemory, ErrorKind};
use std::slice::from_mut;


pub struct Buffer<T> {
    mem : Vec<T>
}

impl<T> Buffer<T> {
    pub fn new(size : usize) -> Buffer<T> {
        Buffer {mem : Vec::with_capacity(size+1)}
    }
    #[inline]
    pub fn chk_cap(v : &Vec<T>) -> bool { v.len() < v.capacity()}

    pub fn push(&mut self,data : T) -> Result<()> {
        if !Buffer::chk_cap(&self.mem) {
            return Err(Error::from(ErrorKind::OutOfMemory));
        }
        self.mem.push(data);
        Ok(())
    }
    pub fn push_from_slice(&mut self,sli : &[T]) {
        self.mem.copy_from_slice(sli);
    }
    pub fn clear_mem(&mut self) {self.mem.clear()}
    pub fn copy_mem<'b>(&self) -> &'b [T] { self.mem.clone().as_slice() }
}