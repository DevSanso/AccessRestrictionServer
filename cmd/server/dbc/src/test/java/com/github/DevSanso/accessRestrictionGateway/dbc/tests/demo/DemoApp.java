package com.github.DevSanso.accessRestrictionGateway.dbc.tests.demo;


import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.domain.EntityScan;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.PropertySource;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;

@SpringBootApplication
@ComponentScan(basePackages={"com.github.DevSanso.accessRestrictionGateway.dbc.*"})
@EnableJpaRepositories(basePackages={"com.github.DevSanso.accessRestrictionGateway.dbc.*"})
@PropertySource("classpath:application.properties")
@EntityScan("com.github.DevSanso.accessRestrictionGateway.dbc.*")
public class DemoApp { }
