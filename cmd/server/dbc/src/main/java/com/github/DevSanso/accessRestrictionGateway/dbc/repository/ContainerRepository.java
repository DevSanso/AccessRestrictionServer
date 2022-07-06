package com.github.DevSanso.accessRestrictionGateway.dbc.repository;


import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;


import com.github.DevSanso.accessRestrictionGateway.dbc.entity.ContainerEntity;

import java.util.List;

@Repository
public interface ContainerRepository extends JpaRepository<ContainerEntity,Long> {
    @Query("SELECT process_name FROM container as c WHERE c.level = :level")
    List<String> processNameFromLevel(@Param("level") Long level);

    @Override
    <S extends ContainerEntity> S saveAndFlush(S entity);
}
