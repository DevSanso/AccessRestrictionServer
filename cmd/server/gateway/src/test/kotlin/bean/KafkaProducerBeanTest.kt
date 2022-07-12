package com.github.DevSanso.accessRestrictionGateway.gateway.configure.bean


import com.github.DevSanso.accessRestrictionGateway.gateway.GatewayApplicationTests
import com.github.DevSanso.accessRestrictionGateway.gateway.configure.KafkaConfigure
import org.apache.kafka.clients.consumer.Consumer
import org.apache.kafka.clients.consumer.ConsumerConfig
import org.apache.kafka.clients.consumer.KafkaConsumer
import org.apache.kafka.clients.producer.KafkaProducer
import org.apache.kafka.clients.producer.ProducerRecord
import org.junit.Assert

import org.junit.Test
import org.junit.runner.RunWith
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.beans.factory.annotation.Qualifier

import org.springframework.boot.test.context.SpringBootTest
import org.springframework.context.annotation.Import
import org.springframework.test.context.junit4.SpringRunner
import java.time.Duration
import java.util.*


@RunWith(SpringRunner::class)
@SpringBootTest(classes = [GatewayApplicationTests::class])
@Import(value = [KafkaProducerBeanConfigure::class])
class KafkaProducerBeanTest {
    @Autowired
    @Qualifier("kafkaProducer")
    lateinit var producer: KafkaProducer<String,ByteArray>

    @Autowired
    lateinit var kafkaConfigure: KafkaConfigure
    @Test
    fun producerTest() {
        val byteArray = "hello world".toByteArray()
        val record = ProducerRecord(kafkaConfigure.thisTopicName,0,"",byteArray)
        val future = producer.send(record)


        val consumer = createConsumer()
        future.get()
        val datas = consumer.poll(Duration.ofMillis(10000))


        val data = datas.first()


        val isMatchData = String(data.value()) == String(byteArray)
        Assert.assertTrue("${String(data.value())} != ${String(byteArray)}",isMatchData)
    }


    fun createConsumer(): Consumer<String, ByteArray> {
        val props = Properties()
        props[ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG] = kafkaConfigure.address
        props[ConsumerConfig.GROUP_ID_CONFIG] = kafkaConfigure.consumerGroupId
        props["key.deserializer"] = "org.apache.kafka.common.serialization.StringDeserializer"
        props["value.deserializer"] = "org.apache.kafka.common.serialization.ByteArrayDeserializer"
        props["auto.offset.reset"] = "latest"
        // Create the consumer using props.
        val consumer: Consumer<String, ByteArray> = KafkaConsumer(props)
        // Subscribe to the topic.
        consumer.subscribe(Collections.singletonList(kafkaConfigure.thisTopicName))
        return consumer
    }
}