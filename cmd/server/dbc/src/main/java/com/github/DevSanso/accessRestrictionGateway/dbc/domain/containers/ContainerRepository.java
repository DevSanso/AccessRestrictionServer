package com.github.DevSanso.accessRestrictionGateway.dbc.domain.containers;


import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;


import java.util.List;

@Repository
public interface ContainerRepository extends JpaRepository<Container,Long> {
    @Query("SELECT process_name FROM container as c WHERE c.level = :level")
    List<String> processNameFromLevel(@Param("level") Long level);

    @Override
    <S extends Container> S saveAndFlush(S entity);
}
