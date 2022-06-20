use std::time::Duration;
use std::io;
use std::io::ErrorKind;

use kafka::client::KafkaClient;
use kafka::producer::{AsBytes, Producer, Record, RequiredAcks};

use accessRestrictionCore::proto::MSAMessage::MSAMessage;
use accessRestrictionCore::proto::MSAHeader::msaheader::State;


macro_rules! build {
    ($b:ident) => {
        (move || {
            return $b.with_ack_timeout(Duration::from_secs(1))
                .with_required_acks(RequiredAcks::One)
                .create()
                .unwrap();
        })()
    }
}

pub struct ProducerProc {
    raw : Producer
}


impl ProducerProc {
    pub fn new(client : KafkaClient) -> Self {
        ProducerProc{raw : build!(Producer::from_client(client))}
    }

    pub async fn send(&mut self,topic : String, message : MSAMessage) -> io::Result<()>{
        let bytes= message.write_to_bytes();
        if bytes.is_err() {
            return Err(io::Error::from(ErrorKind::InvalidInput));
        }
        let message_byte = bytes.unwrap().as_bytes();
        let send_res =self.raw.send(&Record::from_value(topic.as_str(),message_byte));

        if send_res.is_err() {
            return Err(io::Error::from(ErrorKind::ConnectionRefused));
        }


        Ok(())
    }
}
