package com.github.DevSanso.accessRestrictionGateway.gateway.configure

import org.springframework.beans.factory.annotation.Value
import org.springframework.boot.context.properties.ConfigurationProperties
import org.springframework.context.annotation.Configuration
import org.springframework.context.annotation.PropertySource
import org.springframework.stereotype.Component


@Configuration
@ConfigurationProperties(prefix = "kafka")
class KafkaConfigure {
    lateinit var address : String
    lateinit var thisTopicName : String
    lateinit var consumerGroupId : String
}