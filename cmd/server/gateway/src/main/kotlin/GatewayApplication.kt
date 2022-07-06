package com.github.DevSanso.accessRestrictionGateway.gateway

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.context.annotation.ComponentScan


@ComponentScan(basePackages = ["com.github.DevSanso.accessRestrictionGateway"])
@SpringBootApplication
class GatewayApplication

fun main(args: Array<String>) {
	runApplication<GatewayApplication>(*args)
}
