use kafka::client::KafkaClient;
use kafka::consumer::{Consumer, FetchOffset, GroupOffsetStorage,Message};




macro_rules! build {
    ($b:ident,$topic:expr) => {
        (move || {
            return $b.with_topic($topic, $parts)
              .with_fallback_offset(FetchOffset::Earliest)
              .with_offset_storage(GroupOffsetStorage::Kafka)
              .create()
              .unwrap();
        })()
    }
}

pub struct ConsumerProc {
    raw : Consumer
}


impl ConsumerProc {
    pub fn new(client : KafkaClient,topic :&'static str) -> Self {
        ConsumerProc {raw : build!(Consumer::from_client(client),topic)}
    }

    pub fn recv_message(&mut self) -> Vec<Message>  {
        let iter = self.raw.poll().unwrap().iter();
        let map = iter.map(|value|  { return value.messages(); });
        let res = map.fold(Vec::new(),|mut acc, x| {
            acc.copy_from_slice(x);
            acc
        });

        res
    }
}