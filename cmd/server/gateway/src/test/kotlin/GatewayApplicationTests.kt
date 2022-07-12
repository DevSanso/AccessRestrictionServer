package com.github.DevSanso.accessRestrictionGateway.gateway

import com.github.DevSanso.accessRestrictionGateway.gateway.configure.KafkaConfigure
import org.junit.jupiter.api.Test
import org.springframework.boot.context.properties.EnableConfigurationProperties
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.context.annotation.PropertySource

@SpringBootTest
@PropertySource("classpath:application.properties")
@EnableConfigurationProperties(KafkaConfigure::class)
class GatewayApplicationTests {}
