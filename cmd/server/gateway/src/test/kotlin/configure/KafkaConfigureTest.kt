package com.github.DevSanso.accessRestrictionGateway.gateway.configure


import com.github.DevSanso.accessRestrictionGateway.gateway.GatewayApplicationTests
import org.junit.Assert
import org.junit.Test
import org.junit.runner.RunWith
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.test.context.junit4.SpringRunner


@RunWith(SpringRunner::class)
@SpringBootTest(classes = [GatewayApplicationTests::class])
class KafkaConfigureTest {

    @Autowired
    lateinit var kafkaConfigure: KafkaConfigure

    @Test
    fun checkConfigure() {
        val topic = kafkaConfigure.thisTopicName
        val address = kafkaConfigure.address
        Assert.assertTrue(topic,topic.compareTo("gateway") == 0)
        Assert.assertTrue("", address == "")
    }
}