package com.github.DevSanso.accessRestrictionGateway.gateway.configure.bean

import com.github.DevSanso.accessRestrictionGateway.gateway.configure.KafkaConfigure
import org.apache.kafka.clients.producer.KafkaProducer
import org.apache.kafka.clients.producer.Producer
import org.apache.kafka.streams.StreamsConfig
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration

import java.util.*
import javax.annotation.Resource



@Configuration
class KafkaProducerBeanConfigure {
    @Resource
    lateinit var kafkaConfigure: KafkaConfigure

    @Bean("kafkaProducer")
    fun kafkaProducer(): Producer<String, ByteArray> {
        val props = Properties()
        props[StreamsConfig.BOOTSTRAP_SERVERS_CONFIG] = kafkaConfigure.address
        props["key.serializer"] = "org.apache.kafka.common.serialization.StringSerializer"
        props["value.serializer"] = "org.apache.kafka.common.serialization.ByteArraySerializer"
        return KafkaProducer(props)
    }
}