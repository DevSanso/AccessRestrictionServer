package com.github.DevSanso.accessRestrictionGateway.dbc.tests.domain.containers;

import com.github.DevSanso.accessRestrictionGateway.dbc.domain.containers.Container;
import com.github.DevSanso.accessRestrictionGateway.dbc.domain.containers.ContainerRepository;
import com.github.DevSanso.accessRestrictionGateway.dbc.tests.demo.DemoApp;
import org.junit.After;
import org.junit.Test;
import org.junit.Assert;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;

import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import javax.annotation.Resource;
import java.util.Optional;


@RunWith(SpringRunner.class)
@SpringBootTest(classes = DemoApp.class)

public class ContainerRepositoryTest {
    @Autowired
    ContainerRepository containerRepository;

    @After
    public void cleanup() {
        containerRepository.deleteAll();
    }

    @Test
    public void loadContainer() {
        var data = Container.builder().level(0).processName("url").build();
        containerRepository.save(data);
        var tuple = containerRepository.findAll();
        var container = tuple.get(0);
        var level = container.getId();

        Assert.assertEquals(data.level,container.getLevel());
        Assert.assertEquals(data.processName,container.getProcessName());
    }

    @Test
    public void loadContainerFromLevel() {
        containerRepository.save(Container.builder().level(1).processName("url").build());
        var list = containerRepository.processNameFromLevel(1L);
        var name = list.get(0);

        Assert.assertEquals("url",name);
    }

}
