use std::sync::mpsc::Receiver;
use std::io::Result;

use crate::buffer::Buffer;


pub struct BufferAdapter<T> {
    buffers : [Buffer<T>;2],
    w_index : usize,
    r_index : usize
}


impl<T> BufferAdapter<T> {
    pub fn new(buffers : [Buffer<T>;2]) -> BufferAdapter<T> {
        BufferAdapter{buffers,w_index : 0,r_index : 1}
    }
    #[inline]
    fn swap(w_index : usize,r_index : usize) -> (usize,usize) {(r_index, w_index)}

    fn w_buffer<'a>(&mut self) -> &'a mut Buffer<T> {&mut self.buffers[self.w_index]}
    fn r_buffer<'a>(&self) -> &'a Buffer<T> {&self.buffer[self.r_index]}

    pub fn write_buffer(&mut self,data : T) -> Result<()> { self.w_buffer().push(data) }
    pub fn write_buffer_from_slice(&mut self,sli : &[T]) {self.w_buffer().push_from_slice(sli)}
    pub fn read_buffer<'a>(&self) -> &'a [T] {self.r_buffer().copy_mem()}


    pub fn switch(&mut self) {
        let (w_index,r_index) = BufferAdapter::swap(self.w_index,self.r_index);
        self.w_index = w_index;
        self.r_index = r_index;
    }
}


