package com.github.DevSanso.accessRestrictionGateway.dbc.repository;


import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;


import com.github.DevSanso.accessRestrictionGateway.dbc.entity.UrlEntity;

@Repository
public interface UrlRepository extends JpaRepository<UrlEntity,String> {
}
