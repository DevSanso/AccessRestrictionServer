package com.github.DevSanso.accessRestrictionGateway.gateway.configure.bean


import com.github.DevSanso.accessRestrictionGateway.gateway.GatewayApplicationTests
import org.apache.kafka.clients.producer.KafkaProducer
import org.apache.kafka.clients.producer.ProducerRecord
import org.junit.Test
import org.junit.runner.RunWith
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.beans.factory.annotation.Qualifier

import org.springframework.boot.test.context.SpringBootTest
import org.springframework.context.annotation.Import
import org.springframework.test.context.junit4.SpringRunner


@RunWith(SpringRunner::class)
@SpringBootTest(classes = [GatewayApplicationTests::class])
@Import(value = [KafkaProducerBeanConfigure::class])
class KafkaProducerBeanTest {
    @Autowired
    @Qualifier("kafkaProducer")
    lateinit var producer: KafkaProducer<String,ByteArray>
    @Test
    fun producerTest() {

    }
}