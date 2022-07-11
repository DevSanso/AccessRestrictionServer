package com.github.DevSanso.accessRestrictionGateway.dbc.domain.containers;


import com.github.DevSanso.accessRestrictionGateway.dbc.domain.containers.Container;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;


import java.util.List;

@Repository
public interface ContainerRepository extends JpaRepository<Container,Long> {
    @Query(value = "SELECT process_name FROM container  WHERE level = :level",nativeQuery = true)
    List<String> processNameFromLevel(@Param("level") Long level);

}
