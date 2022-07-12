package com.github.DevSanso.accessRestrictionGateway.gateway.configure.bean


import org.apache.kafka.streams.StreamsBuilder
import org.apache.kafka.streams.StreamsConfig
import java.util.*

import com.github.DevSanso.accessRestrictionGateway.gateway.configure.KafkaConfigure
import org.apache.kafka.streams.kstream.KStream
import org.springframework.context.annotation.Bean
import org.springframework.stereotype.Component


import javax.annotation.Resource

@Component
class KafkaInputStreamBeanConfigure {
    @Resource
    lateinit var kafkaConfigure: KafkaConfigure

    @Bean(name = ["kafkaInputStream"])
    fun kafkaInputStream(): KStream<Unit, ByteArray> {
        var props = Properties()
        props[StreamsConfig.BOOTSTRAP_SERVERS_CONFIG] = kafkaConfigure.address

        val builder = StreamsBuilder()
        val s = builder.stream<Unit, ByteArray>(kafkaConfigure.thisTopicName)
        return builder.stream(kafkaConfigure.thisTopicName)
    }
}


