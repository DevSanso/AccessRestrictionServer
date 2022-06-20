use kafka::client::KafkaClient;
use crate::env::KAFKA_HOST;

mod consumer;
mod producer;

pub use consumer::ConsumerProc;
pub use producer::ProducerProc;

pub fn gen_client() -> KafkaClient {
    KafkaClient::new(vec!(KAFKA_HOST.clone().to_owned()))
}

